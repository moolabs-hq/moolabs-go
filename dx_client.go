// Unified Moolabs SDK facade — capability-based public surface (Go).
//
// Go counterpart of sdks/dx/python/moolabs/_dx_client.py. Cross-language
// parity (Task H) asserts identical capability list across py/ts/go.
//
// Usage:
//
//	import "github.com/moolabs/moolabs-go"
//
//	client, err := moolabs.NewMoolabs(moolabs.Config{APIKey: "moo_live_..."})
//	if err != nil { /* handle */ }
//	defer client.Close()
//
//	// Convention-derived hosts; methods promoted from embedded backing services
//	invoices, _, err := client.Billing.ListInvoices(ctx).Execute()
//	wallet, _, err  := client.Wallets.CreateWallet(ctx).WalletCreate(in).Execute()
//
//	// Event-ingest hot path — F2 fallback + G5 buffer
//	delivered, err := client.Usage.IngestEvents(ctx, events)
//
// Constructor changes from rev-1 (pre-2026-05-15 surface):
//   - ClsBaseURL / MeterBaseURL REMOVED — convention-based subdomain derivation
//   - BaseURL is the ROOT DOMAIN (default "moolabs.com")
//   - Buffer / BufferMax fields control G5 in-memory queue (default on, 10k)
//
// 11 capability fields replace the rev-1 Cls / Meter structs. Struct embedding
// promotes methods from backing *<X>APIService values onto the capability
// struct, so customer calls look flat: client.Billing.ListInvoices(...).
// O4 verified zero method-name collisions across multi-class capabilities;
// if a future Arc spec introduces one, the Go compiler will surface it as
// an ambiguous-selector error at customer call site — that's the right time
// to apply the sub-accessor pattern.

package moolabs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// CloudEvents batch ingest path + content-type (CloudEvents 1.0 / Meter spec).
// We POST directly via net/http here instead of the auto-generated
// EventsAPI.IngestEvents() builder for two reasons:
//
//  1. The Meter OpenAPI spec declares TWO content-types for /api/v1/events:
//     application/cloudevents+json (single Event body) and
//     application/cloudevents-batch+json (array body). openapi-generator's
//     Go template binds the SINGLE variant — `.Event(event)` takes one
//     Event, not a batch. Bulk ingest with the generated builder would
//     mean N round-trips per call, defeating the buffer's batching purpose.
//  2. The F2 fallback chain needs to swap base URLs per-attempt; rebuilding
//     a full APIClient per attempt (which the generated builder requires)
//     is heavier than a raw POST that just substitutes the host.
const (
	ingestPath             = "/api/v1/events"
	ingestBatchContentType = "application/cloudevents-batch+json"
)

// IngestError represents an HTTP failure of an ingest POST. Use Terminal()
// to decide whether to retry/buffer or surface to the caller.
//
// Terminal codes (401/403, 400/422, 404): retry will fail identically with
// the same key/body/URL. Buffering or re-walking the F2 chain on a terminal
// error silently loses events with no error visible to the caller, and
// corrupts the F2 resolver's recently_failed set on an innocent URL. Always
// surface terminal errors to the customer immediately.
//
// Transient codes (5xx, 408, 429) and network errors: retry-friendly. The
// F2 chain re-walk and G5 buffer enqueue paths apply.
type IngestError struct {
	StatusCode int    // 0 if the request never reached a response (network error)
	URL        string // the ingest URL we POSTed to
	Body       string // truncated response body (max 4 KB)
	Cause      error  // underlying transport error if StatusCode == 0
}

func (e *IngestError) Error() string {
	if e.StatusCode == 0 {
		return fmt.Sprintf("moolabs: posting events to %s: %v", e.URL, e.Cause)
	}
	return fmt.Sprintf("moolabs: ingest POST %s returned %d: %s", e.URL, e.StatusCode, e.Body)
}

func (e *IngestError) Unwrap() error { return e.Cause }

// Terminal reports whether this error is non-retryable.
func (e *IngestError) Terminal() bool {
	switch e.StatusCode {
	case 400, 401, 403, 404, 422:
		return true
	}
	return false
}

// isTerminalIngestError peeks through error wrapping to find an IngestError
// and return its Terminal() classification. Returns false for non-IngestError
// types (treated as transient/retryable by default).
func isTerminalIngestError(err error) bool {
	var ie *IngestError
	if errors.As(err, &ie) {
		return ie.Terminal()
	}
	return false
}

// postEventsBatch sends `events` as a JSON array to <baseURL><ingestPath>
// with the batch content-type and Bearer auth. Returns (delivered, *IngestError)
// where `delivered` is len(events) on 2xx and 0 otherwise. Callers should
// check isTerminalIngestError(err) to decide whether to retry/buffer.
func postEventsBatch(ctx context.Context, baseURL, apiKey string, events []any) (int, error) {
	body, err := json.Marshal(events)
	if err != nil {
		return 0, fmt.Errorf("moolabs: marshalling events: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+ingestPath, bytes.NewReader(body))
	if err != nil {
		return 0, fmt.Errorf("moolabs: building ingest request: %w", err)
	}
	req.Header.Set("Content-Type", ingestBatchContentType)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, &IngestError{URL: baseURL, Cause: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// Drain body for connection reuse; ignore content (spec says 204).
		_, _ = io.Copy(io.Discard, resp.Body)
		return len(events), nil
	}
	// Read and truncate body for diagnostics
	const maxBody = 4096
	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, maxBody))
	return 0, &IngestError{
		StatusCode: resp.StatusCode,
		URL:        baseURL,
		Body:       string(respBody),
	}
}

const defaultBaseURL = "moolabs.com"

// Config configures a Moolabs client.
type Config struct {
	// APIKey is the customer's dashboard-issued Moolabs API key. The same
	// key authenticates against every backend (BFF / Meter / Arc) via the
	// shared key store (C1 / Shape A).
	APIKey string

	// BaseURL is the root domain. Default: "moolabs.com". Self-hosted
	// customers pass their own root (e.g. "moolabs.example.com"); the SDK
	// derives api./meter./arc. subdomains internally.
	BaseURL string

	// Buffer controls the G5 in-memory ingest buffer. When *true (default),
	// F2-chain-exhaustion enqueues events for background retry instead of
	// returning an error. When *false, IngestEvents returns the chain's
	// final error directly. nil means "use default" (= enabled).
	Buffer *bool

	// BufferMax sets the bounded queue size. Default: 1000 (matches
	// DefaultIngestBufferConfig.MaxSize). Customers with higher
	// throughput should raise this to avoid drop_oldest on steady-state
	// load.
	BufferMax int

	// Logger receives the SDK's per-event diagnostic warnings — terminal
	// drops (auth/validation upstream error), buffer overflow, drain
	// failures, abandoned-on-shutdown.
	//
	// Default behavior: NIL = NoopLogger = NO OUTPUT. The library never
	// writes to stderr/console unless the customer explicitly opts in by
	// providing a Logger. This means SDK debugging is available when
	// wanted but stderr stays clean by default.
	//
	// When provided, the Logger receives ONE call per event (not throttled
	// — customers throttle in their own adapter if they care). Per-event
	// data + Stats() counters together give: counters for periodic
	// monitoring (rate, total), logs for per-event context (which status
	// code, which batch size, which error message).
	//
	// History: PR #395 review round 2 (M-NEW-2) flagged that the original
	// log.Printf calls polluted customer stderr. Initial fix was a Logger
	// interface defaulting to a StdlibLogger — still polluted by default.
	// Final design: NoopLogger default. Logger interface present so
	// customers can opt in to verbose SDK diagnostics for debugging.
	//
	// Logger interface + NoopLogger live in dx_logger.go (leaf file, no
	// dependency on generated code, compilable with leaf-only tests).
	Logger Logger
}

// Moolabs is the top-level SDK client. Use NewMoolabs to construct, defer
// Close() to drain the ingest buffer + release resources.
//
// Each capability field embeds the corresponding backing *<X>APIService
// pointers, so methods are promoted onto the capability for clean call
// sites: client.Billing.ListInvoices(...).
type Moolabs struct {
	apiKey         string
	baseURL        string
	bufferEnabled  bool
	bufferMax      int
	logger         Logger
	ingestResolver *IngestURLResolver
	ingestBuffer   *IngestBuffer
	bufferOnce     sync.Once

	bffClient   *APIClient
	meterClient *APIClient
	arcClient   *APIClient
	// acuteClient — sdk-cost-capability-acute-backing US-008 (Ralph iter 23 follow-up).
	// Backs the cost capability's direct ACUTE wiring (CostEventsAPI + SdkIngestAPI).
	acuteClient *APIClient

	// 11 capability fields. Each embeds backing services so methods promote.
	Usage         *UsageNamespace
	Customers     *CustomersNamespace
	Catalog       *CatalogNamespace
	Subscriptions *SubscriptionsNamespace
	Entitlements  *EntitlementsNamespace
	Wallets       *WalletsNamespace
	Credits       *CreditsNamespace
	Billing       *BillingNamespace
	Collections   *CollectionsNamespace
	Cost          *CostNamespace
	Notifications *NotificationsNamespace
}

// ── Capability namespace types ──
//
// Each namespace embeds the *<X>APIService values for its backing classes
// (see CapabilityMap in dx_routing.go). Methods on those services are
// promoted to the outer struct so customer code is flat.
//
// Special case: UsageNamespace shadows IngestEvents with an F2+G5 routed
// version. Customers who want the raw builder API can reach it via the
// explicit embedded field: client.Usage.EventsAPIService.IngestEvents(ctx).

type CustomersNamespace struct {
	*CustomersAPIService
	*SubjectsAPIService
}

type CatalogNamespace struct {
	*ProductCatalogAPIService
	*RateCardsAPIService
}

type SubscriptionsNamespace struct {
	*MeterSubscriptionsAPIService
}

type EntitlementsNamespace struct {
	*EntitlementsAPIService
}

type WalletsNamespace struct {
	*WalletsAPIService
}

type CreditsNamespace struct {
	*GrantsAPIService
	*LedgerAPIService
	*AutoTopupAPIService
}

type BillingNamespace struct {
	*MeterBillingAPIService
	*RatingAPIService
	*FxRatesAPIService
}

type CollectionsNamespace struct {
	*AccountsAPIService
	*AccountTeamAPIService
	*AnalyticsAPIService
	*CasesAPIService
	*CashCreditsAPIService
	*ArcCommunicationsAPIService // stitcher-renamed tag
	*CreditMemosAPIService
	*DisputesAPIService
	*EscalationsAPIService
	*HandoffsAPIService
	*NotesAPIService
	*PaymentsAPIService
	*PlansAPIService
	*PromisesAPIService
	*RemittancesAPIService
	*ReportsAPIService
	*TasksAPIService
}

// CostNamespace — sdk-cost-capability-acute-backing US-008 (Ralph iter 23
// follow-up). Embeds CostEventsAPIService + SdkIngestAPIService from the
// ACUTE backend (was *AcuteProxyAPIService from BFF before US-008's remap).
type CostNamespace struct {
	*CostEventsAPIService
	*SdkIngestAPIService
}

type NotificationsNamespace struct {
	*NotificationsAPIService
	*AlertsAPIService
}

// UsageNamespace embeds EventsAPIService + MetersAPIService (so ListEvents,
// CreateMeter, QueryMeter etc. work via promotion) AND adds an IngestEvents
// method that shadows the promoted one and routes through the F2+G5 path.
type UsageNamespace struct {
	*EventsAPIService
	*MetersAPIService

	// Hooks injected by NewMoolabs for the F2+G5 path.
	resolver        *IngestURLResolver
	buffer          *IngestBuffer
	apiKey          string
	makeClientAtURL func(host string) *APIClient

	// Producer-channel pattern (post-PR #395 round-5): the customer's
	// IngestEvents call does a non-blocking channel send (~50 ns), then
	// a dedicated producer goroutine moves events from the channel to
	// the buffer. This keeps the customer's goroutine off the buffer's
	// mutex entirely. Channel does drop-newest on overflow (counted in
	// ingestQueueDropped); buffer does drop-oldest. Both are surfaced
	// via Stats().
	// Producer-channel for non-blocking ingest (round-5 extraction).
	// Extracted to dx_producer.go so it's leaf-testable without
	// depending on the openapi-generator-generated EventsAPIService etc.
	// The delegate is u.buffer.Enqueue.
	producer *ProducerChannel
}

// NewMoolabs constructs a Moolabs client. Returns an error on invalid
// APIKey / BaseURL (so customer typos crash at construction, not on first
// capability call).
//
// No network I/O at construction. Per-backend APIClient instances are
// created eagerly (cheap; just builds Configuration). When Config.Buffer
// is true (the default), the G5 buffer worker goroutine starts inside
// wireCapabilities — Close() MUST be called to stop it. F2 discovery is
// deferred until the first IngestEvents call; the producer-channel
// goroutine is also lazy (started on first Submit via sync.Once).
func NewMoolabs(cfg Config) (*Moolabs, error) {
	if cfg.APIKey == "" {
		return nil, errors.New("moolabs: Config.APIKey is required")
	}
	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	// Apex ("moolabs.com") is a marketing/branding host, not an env root —
	// the ALB cert is "*.prod.moolabs.com" only. Rewrite the customer-
	// supplied baseURL ONCE at construction so all downstream subdomain
	// composition (DeriveHost, IngestURLResolver) sees the effective
	// env-rooted host. Explicit env roots and self-hosted bases pass
	// through unchanged. See dx_urls.go:ResolveEffectiveBaseURL.
	resolvedBaseURL, err := ResolveEffectiveBaseURL(baseURL, cfg.APIKey)
	if err != nil {
		return nil, fmt.Errorf("moolabs: BaseURL invalid: %w", err)
	}
	baseURL = resolvedBaseURL
	// Validate BaseURL early — fail fast on typos.
	for backend := range SubdomainMap {
		if _, err := DeriveHost(backend, baseURL); err != nil {
			return nil, fmt.Errorf("moolabs: BaseURL invalid: %w", err)
		}
	}

	bufferEnabled := true
	if cfg.Buffer != nil {
		bufferEnabled = *cfg.Buffer
	}
	bufferMax := cfg.BufferMax
	if bufferMax == 0 {
		bufferMax = 1000   // matches DefaultIngestBufferConfig.MaxSize
	}

	logger := cfg.Logger
	if logger == nil {
		logger = NoopLogger{}   // default = no output (library never pollutes customer stderr)
	}

	m := &Moolabs{
		apiKey:        cfg.APIKey,
		baseURL:       baseURL,
		bufferEnabled: bufferEnabled,
		bufferMax:     bufferMax,
		logger:        logger,
	}

	// Per-backend APIClient instances. Errors are ignored on the DeriveHost
	// calls below since we already validated all three above.
	bffHost, _ := DeriveHost(BackendBFF, baseURL)
	meterHost, _ := DeriveHost(BackendMeter, baseURL)
	arcHost, _ := DeriveHost(BackendArc, baseURL)
	// ACUTE — sdk-cost-capability-acute-backing US-008 (Ralph iter 23 follow-up).
	// Cost capability methods (CostEventsAPI + SdkIngestAPI) route through
	// acute.{baseURL} directly, not through the BFF cost-ingest-proxy.
	acuteHost, _ := DeriveHost(BackendAcute, baseURL)
	m.bffClient = m.makeClientAtURL(bffHost)
	m.meterClient = m.makeClientAtURL(meterHost)
	m.arcClient = m.makeClientAtURL(arcHost)
	m.acuteClient = m.makeClientAtURL(acuteHost)

	// F2 resolver — wired but not consulted until first IngestEvents call.
	resolver, err := NewIngestURLResolver(baseURL, m.discoverTenantConfig)
	if err != nil {
		return nil, fmt.Errorf("moolabs: building IngestURLResolver: %w", err)
	}
	m.ingestResolver = resolver

	m.wireCapabilities()
	return m, nil
}

// Close drains the ingest buffer (if started), bounded by the buffer's
// ShutdownFlushTimeout, and releases per-backend resources. Idempotent.
func (m *Moolabs) Close() error {
	// Stop the producer first so it drains pending channel events into
	// the buffer; then close the buffer so it can do its final HTTP drain.
	if m.Usage != nil {
		m.Usage.stopProducer()
	}
	if m.ingestBuffer != nil {
		return m.ingestBuffer.Close()
	}
	return nil
}

// IngestQueueDropped returns events dropped at the producer-channel
// layer (channel full at the moment of Submit). Distinct from
// IngestBufferStats.Dropped which counts buffer-layer drop_oldest after
// the producer moved events through. Delegated to ProducerChannel.
func (u *UsageNamespace) IngestQueueDropped() int64 {
	if u.producer == nil {
		return 0
	}
	return u.producer.IngestQueueDropped()
}

// ProducerPanics returns how often the producer goroutine recovered
// from a panic and respawned. Non-zero indicates a buggy delegate
// (typically a customer Logger that panicked). Delegated to
// ProducerChannel.
func (u *UsageNamespace) ProducerPanics() int64 {
	if u.producer == nil {
		return 0
	}
	return u.producer.ProducerPanics()
}

// ── F2+G5 wiring on UsageNamespace ──

// IngestEvents POSTs events via the F2 fallback chain. On chain exhaustion
// with Buffer enabled, enqueues events to the G5 in-memory buffer and
// returns (0, nil) to indicate "queued, not delivered yet". With Buffer
// disabled, returns the underlying error.
//
// Shadows the IngestEvents method promoted from *EventsAPIService. Customers
// who need the raw builder (e.g. for per-call header overrides) can reach
// the promoted one via client.Usage.EventsAPIService.IngestEvents(ctx).
// IngestEvents posts events.
//
// **Default mode (Buffer enabled):** non-blocking. Customer's
// goroutine does ONLY a non-blocking channel send (~50 ns) — never
// the buffer's mutex, never the network. A producer goroutine drains
// the channel into the buffer; a drain goroutine pulls from the
// buffer and POSTs via the F2 chain. The customer's goroutine is
// off the SDK after the channel send.
//
// Returns (len(events), nil) on successful channel send. On channel
// full (overflow), increments Stats().IngestQueueDropped and returns
// (0, nil) — the call still appears successful from the customer's
// perspective; the drop signal is observable via Stats() polling.
//
// **Strict-sync mode (Config.Buffer = &false):** blocking. HTTP POST
// runs inline on the caller's goroutine; HTTP errors propagate up.
// Returns (delivered, error). Use only when the caller specifically
// needs delivery confirmation per call.
//
// Design history:
//   - Pre-round-4: buffer was failure-only; every call blocked for
//     full HTTP round-trip even on success.
//   - Round-4: direct enqueue on customer goroutine; ~200ns-10µs
//     under mutex.
//   - Round-5 (this commit): producer-channel pattern; ~50 ns
//     channel send. Customer goroutine never touches the buffer's
//     mutex.
func (u *UsageNamespace) IngestEvents(ctx context.Context, events []any) (int, error) {
	if len(events) == 0 {
		return 0, nil
	}

	if u.buffer != nil {
		// Non-blocking channel send via ProducerChannel. SubmitAndCount
		// returns len(events) when the producer accepted the batch, 0
		// when the channel was full (events dropped at the producer
		// queue; counter bumped — observable via IngestQueueDropped()).
		//
		// Codex round-6 HIGH-2: the helper formalizes the count contract
		// so it can't drift. Pre-fix, IngestEvents discarded Submit's
		// boolean and always returned len(events), silently lying to the
		// caller on drop. Tests pin SubmitAndCount in dx_producer_test.go.
		return u.producer.SubmitAndCount(events), nil
	}

	// Strict-sync mode: caller wants delivery confirmation per call.
	url := u.resolver.GetIngestURL()
	delivered, err := postEventsBatch(ctx, url, u.apiKey, events)
	if err != nil {
		if isTerminalIngestError(err) {
			return 0, err
		}
		u.resolver.ReportPostOutcome(url, false)
		return 0, fmt.Errorf("moolabs: ingest failed and buffer disabled: %w", err)
	}
	u.resolver.ReportPostOutcome(url, true)
	return delivered, nil
}

// stopProducer signals the producer to exit + waits for drain. Called
// from Moolabs.Close(). Delegated to ProducerChannel.Stop().
func (u *UsageNamespace) stopProducer() {
	if u.producer == nil {
		return // never constructed
	}
	u.producer.Stop()
}

// ── internals ──

func (m *Moolabs) makeClientAtURL(host string) *APIClient {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: host}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+m.apiKey)
	return NewAPIClient(cfg)
}

func (m *Moolabs) wireCapabilities() {
	m.Customers = &CustomersNamespace{
		CustomersAPIService: m.meterClient.CustomersAPI,
		SubjectsAPIService:  m.meterClient.SubjectsAPI,
	}
	m.Catalog = &CatalogNamespace{
		ProductCatalogAPIService: m.meterClient.ProductCatalogAPI,
		RateCardsAPIService:      m.bffClient.RateCardsAPI,
	}
	m.Subscriptions = &SubscriptionsNamespace{
		MeterSubscriptionsAPIService: m.meterClient.MeterSubscriptionsAPI,
	}
	m.Entitlements = &EntitlementsNamespace{
		EntitlementsAPIService: m.meterClient.EntitlementsAPI,
	}
	m.Wallets = &WalletsNamespace{
		WalletsAPIService: m.bffClient.WalletsAPI,
	}
	m.Credits = &CreditsNamespace{
		GrantsAPIService:    m.bffClient.GrantsAPI,
		LedgerAPIService:    m.bffClient.LedgerAPI,
		AutoTopupAPIService: m.bffClient.AutoTopupAPI,
	}
	m.Billing = &BillingNamespace{
		MeterBillingAPIService: m.meterClient.MeterBillingAPI,
		RatingAPIService:       m.bffClient.RatingAPI,
		FxRatesAPIService:      m.bffClient.FxRatesAPI,
	}
	m.Collections = &CollectionsNamespace{
		AccountsAPIService:          m.arcClient.AccountsAPI,
		AccountTeamAPIService:       m.arcClient.AccountTeamAPI,
		AnalyticsAPIService:         m.arcClient.AnalyticsAPI,
		CasesAPIService:             m.arcClient.CasesAPI,
		CashCreditsAPIService:       m.arcClient.CashCreditsAPI,
		ArcCommunicationsAPIService: m.arcClient.ArcCommunicationsAPI,
		CreditMemosAPIService:       m.arcClient.CreditMemosAPI,
		DisputesAPIService:          m.arcClient.DisputesAPI,
		EscalationsAPIService:       m.arcClient.EscalationsAPI,
		HandoffsAPIService:          m.arcClient.HandoffsAPI,
		NotesAPIService:             m.arcClient.NotesAPI,
		PaymentsAPIService:          m.arcClient.PaymentsAPI,
		PlansAPIService:             m.arcClient.PlansAPI,
		PromisesAPIService:          m.arcClient.PromisesAPI,
		RemittancesAPIService:       m.arcClient.RemittancesAPI,
		ReportsAPIService:           m.arcClient.ReportsAPI,
		TasksAPIService:             m.arcClient.TasksAPI,
	}
	m.Cost = &CostNamespace{
		// sdk-cost-capability-acute-backing US-008 (Ralph iter 23 follow-up):
		// route client.cost.* DIRECT to acute.{baseURL} instead of the BFF
		// cost-ingest-proxy. Two backing services per HLD Appendix D.5 / RFC
		// + the parity-check'd dx_routing.go CapabilityMap["cost"] entry.
		CostEventsAPIService: m.acuteClient.CostEventsAPI,
		SdkIngestAPIService:  m.acuteClient.SdkIngestAPI,
	}
	m.Notifications = &NotificationsNamespace{
		NotificationsAPIService: m.meterClient.NotificationsAPI,
		AlertsAPIService:        m.bffClient.AlertsAPI,
	}
	buf := m.lazyBuffer()
	usage := &UsageNamespace{
		EventsAPIService: m.meterClient.EventsAPI,
		MetersAPIService: m.meterClient.MetersAPI,
		resolver:         m.ingestResolver,
		buffer:           buf,
		apiKey:           m.apiKey,
		makeClientAtURL:  m.makeClientAtURL,
	}
	// Producer-channel only wired when buffer is enabled. Channel cap
	// scales with BufferMax: max(1024, BufferMax/8) per round-4 I-NEW-5.
	if buf != nil {
		chanCap := 1024
		if m.bufferMax/8 > chanCap {
			chanCap = m.bufferMax / 8
		}
		usage.producer = NewProducerChannel(chanCap, func(events []any) {
			_ = buf.Enqueue(events)
		}, m.logger)
	}
	m.Usage = usage
}

func (m *Moolabs) lazyBuffer() *IngestBuffer {
	if !m.bufferEnabled {
		return nil
	}
	m.bufferOnce.Do(func() {
		cfg := DefaultIngestBufferConfig
		cfg.MaxSize = m.bufferMax
		cfg.Logger = m.logger    // propagate customer-supplied logger (NoopLogger if none)
		buf, err := NewIngestBuffer(m.bufferDrain, cfg)
		if err != nil {
			// Validation failure on default config is a programmer error;
			// surface loudly.
			panic("moolabs: building IngestBuffer: " + err.Error())
		}
		// Publish the pointer BEFORE starting the worker. The worker
		// calls m.bufferDrain, which reads m.ingestBuffer for terminal-
		// drop counter bumps — if Start ran first the worker could
		// observe nil on the first tick (M-NEW-8, round-4 review).
		m.ingestBuffer = buf
		buf.Start()
	})
	return m.ingestBuffer
}

// bufferDrain is the IngestBuffer's drain callback. It retries delivery
// via the F2 chain using the same direct-POST helper as IngestEvents,
// avoiding the single-event auto-generated builder.
//
// Terminal errors (401/403/400/422/404) DISCARD the batch — retry will
// fail identically with the same key/body. Returning len(events) tells
// the buffer the events are "delivered" (removed from queue). A WARN
// log records the drop so operators see the failure mode. Without this,
// a single bad API key fills the buffer forever and silently loses
// every customer event.
func (m *Moolabs) bufferDrain(ctx context.Context, events []any) (int, error) {
	url := m.ingestResolver.GetIngestURL()
	delivered, err := postEventsBatch(ctx, url, m.apiKey, events)
	if err != nil {
		if isTerminalIngestError(err) {
			// Bump counter (always; cheap and aggregable via Stats()) AND
			// emit a per-event log (only if customer provided a Logger).
			// Counter is for periodic monitoring; log is for debugging the
			// specific failure (which status, which batch size, raw error).
			if m.ingestBuffer != nil {
				m.ingestBuffer.recordTerminalDrop(len(events))
			}
			var ie *IngestError
			_ = errors.As(err, &ie)
			status := 0
			if ie != nil {
				status = ie.StatusCode
			}
			m.logger.Warn("moolabs.ingest_buffer.terminal_drop",
				"status", status,
				"count", len(events),
				"err", err,
			)
			return len(events), nil
		}
		m.ingestResolver.ReportPostOutcome(url, false)
		return 0, err
	}
	m.ingestResolver.ReportPostOutcome(url, true)
	return delivered, nil
}

// discoverTenantConfig is the F2 resolver's DiscoveryFn. Calls GET
// /tenant/config on the BFF via raw http.Client (no dependency on the
// generated APIClient for this call — keeps the resolver decoupled from
// any specific service).
func (m *Moolabs) discoverTenantConfig() (map[string]interface{}, error) {
	host, err := DeriveHost(BackendBFF, m.baseURL)
	if err != nil {
		return nil, err
	}
	url := host + DiscoveryPath
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+m.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("moolabs: /tenant/config returned %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("moolabs: /tenant/config parse: %w", err)
	}
	return out, nil
}
