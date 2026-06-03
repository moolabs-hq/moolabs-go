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
	return isTerminalStatusCode(e.StatusCode)
}

// isTerminalStatusCode reports whether an HTTP status is non-retryable for
// ingest (auth/validation/not-found): retrying with the same key/body fails
// identically. Shared by the usage IngestError path and the cost-buffer drain.
func isTerminalStatusCode(code int) bool {
	switch code {
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
//
// Path-doubling guard (sibling of Python _strip_path / TS stripPath):
// the F2 IngestURLResolver emits full URLs INCLUDING /api/v1/events, but
// this function appends ingestPath. Strip any existing path from baseURL
// first so the concatenation can't double. Verified by the 2026-06-02
// dev.moolabs.com live test.
func postEventsBatch(ctx context.Context, baseURL, apiKey string, events []any) (int, error) {
	body, err := json.Marshal(events)
	if err != nil {
		return 0, fmt.Errorf("moolabs: marshalling events: %w", err)
	}
	baseURL = stripPath(baseURL)
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

// defaultIngestTimeout bounds each ingest HTTP attempt when Config.Timeout
// is unset. Generous for batch ingest; prevents an unreachable server from
// hanging a caller goroutine or a buffer drain indefinitely.
const defaultIngestTimeout = 30 * time.Second

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

	// Timeout bounds each INGEST HTTP attempt (the background buffer drains)
	// so an unreachable/black-hole server can never hang a drain goroutine
	// indefinitely. Zero means "use default" (= defaultIngestTimeout).
	//
	// Scope: applied per-attempt via context.WithTimeout in the usage + cost
	// buffer drains ONLY. It is deliberately NOT installed as a client-wide
	// HTTP timeout on the generated backend clients — doing so would also cap
	// NON-ingest calls (billing, reports, credits, subscriptions, …) at the
	// ingest ceiling and break legitimately long customer operations. Those
	// calls are bounded by the caller's own context deadline instead.
	Timeout time.Duration

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
	timeout        time.Duration
	logger         Logger
	ingestResolver *IngestURLResolver
	ingestBuffer   *IngestBuffer
	bufferOnce     sync.Once
	// Cost capability gets its own buffer (separate drain endpoint from usage:
	// ACUTE /api/v1/cost/ingest/batch vs Meter /api/v1/events). Same generic
	// IngestBuffer + ProducerChannel machinery as usage.
	costBuffer     *IngestBuffer
	costBufferOnce sync.Once

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

	// Events is the dual-lane unified ingest namespace (US-011). Customers
	// who don't want to know about usage vs cost can call
	// client.Events.Ingest(ctx, IngestArgs{...}) once and let the SDK
	// translate the args into a CloudEvent envelope. The empty-lane guard
	// rejects calls that supply neither lane synchronously at the call
	// site (FR-6). NOT a capability in CapabilityMap — there's no backing
	// API service; the namespace is a thin wrapper over the meter F2 chain.
	Events *EventsNamespace
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

	// G5 buffer + producer-channel for non-blocking cost ingest — mirrors
	// UsageNamespace. The customer's IngestEvents call does a non-blocking
	// channel send (~50 ns); a producer goroutine moves batches into the
	// buffer; the buffer's drain worker POSTs via the generated cost client.
	// Nil when Buffer is disabled (strict-sync mode).
	buffer   *IngestBuffer
	producer *ProducerChannel

	// Meter-routing hooks for the new unified IngestEvent method (US-010
	// of the SDK Unified Ingest Methods PRD). The new ergonomic
	// IngestEvent posts a CloudEvent envelope to the SAME meter endpoint
	// as UsageNamespace.IngestEvent — both lanes are unified at the wire
	// level. These share the same *IngestURLResolver and meter buffer +
	// producer as UsageNamespace, populated by wireCapabilities. They are
	// independent of the acute-routed buffer/producer above (which is
	// still used by the legacy IngestEvents batch method).
	meterResolver *IngestURLResolver
	meterBuffer   *IngestBuffer
	meterProducer *ProducerChannel
	apiKey        string

	// One-shot deprecation latch for the legacy CostNamespace.IngestEvents
	// batch method. Customers calling that method see a single warn log
	// per instance pointing at the new IngestEvent surface (US-010).
	legacyIngestEventsDeprecationOnce sync.Once
	logger                            Logger
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
		bufferMax = 1000 // matches DefaultIngestBufferConfig.MaxSize
	}

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = defaultIngestTimeout
	}

	logger := cfg.Logger
	if logger == nil {
		logger = NoopLogger{} // default = no output (library never pollutes customer stderr)
	}

	m := &Moolabs{
		apiKey:        cfg.APIKey,
		baseURL:       baseURL,
		bufferEnabled: bufferEnabled,
		bufferMax:     bufferMax,
		timeout:       timeout,
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
	// Stop producers first so they drain pending channel events into their
	// buffers; then close the buffers so they can do their final HTTP drain.
	if m.Usage != nil {
		m.Usage.stopProducer()
	}
	if m.Cost != nil {
		m.Cost.stopProducer()
	}
	var firstErr error
	if m.ingestBuffer != nil {
		if err := m.ingestBuffer.Close(); err != nil {
			firstErr = err
		}
	}
	if m.costBuffer != nil {
		if err := m.costBuffer.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
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

// IngestEvent is the ergonomic struct-arg usage-lane ingest method
// (US-009 — SDK Unified Ingest Methods PRD §3.2). Mirror of Python
// UsageNamespace.ingest_event and TypeScript UsageNamespace.ingestEvent.
//
// Builds a CloudEvent envelope from args (see buildEnvelope in
// dx_envelope.go), performs boundary checks synchronously, then routes
// through the existing F2 + G5 path:
//
//   - When Buffer is enabled (default), enqueues via the producer channel
//     (~50 ns) and returns IngestResult{Transport: TransportBuffered}.
//   - When Buffer is disabled (Config.Buffer = &false), POSTs inline via
//     the F2 chain and returns IngestResult{Transport: TransportSync}.
//
// On a boundary-check failure (empty required field, non-finite Value,
// missing span_id, non-JSON-serializable Meta), returns IngestResult{}
// and the error. The envelope is NOT enqueued, so a bad call cannot leak
// half-formed events into the buffer (FR-6).
//
// On HTTP error in strict-sync mode, returns IngestResult{} and the
// wrapped error. Terminal errors (401/403/400/422/404) are surfaced as
// *IngestError so callers can errors.As() and decide retry-or-not.
//
// TenantID is NOT a field on IngestEventArgs (FR-3) — the server
// derives tenant identity from the API key.
func (u *UsageNamespace) IngestEvent(ctx context.Context, args IngestEventArgs) (IngestResult, error) {
	if args.MeterSlug == "" {
		return IngestResult{}, errors.New("meter_slug must be a non-empty string")
	}
	envelope, err := buildEnvelope(buildEnvelopeArgs{
		EventType:  args.EventType,
		CustomerID: args.CustomerID,
		EntityID:   args.EntityID,
		MeterSlug:  args.MeterSlug,
		HasValue:   true,
		Value:      args.Value,
		// Well-known top-level data.* keys (canonical wire shape).
		Provider:          args.Provider,
		Model:             args.Model,
		TotalInputTokens:  args.TotalInputTokens,
		TotalOutputTokens: args.TotalOutputTokens,
		TotalTokens:       args.TotalTokens,
		LatencyMs:         args.LatencyMs,
		Status:            args.Status,
		EventID:           args.EventID,
		Source:            args.Source,
		Time:              args.Time,
		Meta:              args.Meta,
	})
	if err != nil {
		return IngestResult{}, err
	}
	acceptedAt := time.Now().UTC()
	eventID, _ := envelope["id"].(string)

	if u.buffer != nil {
		// Buffered (default) — non-blocking channel send. Producer is
		// guaranteed non-nil whenever buffer is non-nil (see
		// wireCapabilities below). SubmitAndCount's return value is
		// intentionally ignored: an overflow drop is observable via
		// IngestQueueDropped() and we still report the envelope as
		// "buffered" — the customer's contract is "we accepted it; check
		// stats for backpressure", identical to IngestEvents.
		_ = u.producer.SubmitAndCount([]any{envelope})
		return IngestResult{
			EventID:    eventID,
			Transport:  TransportBuffered,
			AcceptedAt: acceptedAt,
		}, nil
	}

	// Strict-sync: caller wants delivery confirmation per call.
	url := u.resolver.GetIngestURL()
	if _, postErr := postEventsBatch(ctx, url, u.apiKey, []any{envelope}); postErr != nil {
		if isTerminalIngestError(postErr) {
			return IngestResult{}, postErr
		}
		u.resolver.ReportPostOutcome(url, false)
		return IngestResult{}, fmt.Errorf("moolabs: ingest failed and buffer disabled: %w", postErr)
	}
	u.resolver.ReportPostOutcome(url, true)
	return IngestResult{
		EventID:    eventID,
		Transport:  TransportSync,
		AcceptedAt: acceptedAt,
	}, nil
}

// stopProducer signals the producer to exit + waits for drain. Called
// from Moolabs.Close(). Delegated to ProducerChannel.Stop().
func (u *UsageNamespace) stopProducer() {
	if u.producer == nil {
		return // never constructed
	}
	u.producer.Stop()
}

// ── G5 buffering on CostNamespace (mirrors UsageNamespace) ──

// IngestEvents enqueues cost events for buffered, non-blocking delivery to the
// ACUTE cost-ingest endpoint, mirroring UsageNamespace.IngestEvents.
//
// Default mode (Buffer enabled): non-blocking — a ~50 ns channel send; a
// producer goroutine moves the batch into the buffer; the buffer's drain
// worker POSTs via the generated CostEventsAPI batch endpoint, holding +
// retrying on transient failure. Returns (len(events), nil) on accept, or
// (0, nil) on producer-channel overflow (observable via IngestQueueDropped()).
//
// Strict-sync mode (Config.Buffer = &false): blocking — POSTs inline via the
// generated client and returns (delivered, error).
//
// Retry is safe: each CostEventIngest carries idempotency_key and the backend
// dedups, so re-sending a held batch cannot double-charge.
func (c *CostNamespace) IngestEvents(ctx context.Context, events []CostEventIngest) (int, error) {
	if len(events) == 0 {
		return 0, nil
	}
	// US-010 deprecation latch: warn ONCE per instance the first time a
	// customer hits the legacy batch path. The unified IngestEvent surface
	// (which posts a CloudEvent envelope to meter, not a batched
	// CostEventIngest list to acute) is the going-forward path.
	if c.logger != nil {
		c.legacyIngestEventsDeprecationOnce.Do(func() {
			c.logger.Warn("moolabs.cost.ingest_events.deprecated",
				"detail", "CostNamespace.IngestEvents is deprecated; use IngestEvent for per-call ingestion or IngestEventsBatch directly for the legacy acute path",
			)
		})
	}
	if c.buffer != nil {
		anyEvents := make([]any, len(events))
		for i := range events {
			anyEvents[i] = events[i]
		}
		return c.producer.SubmitAndCount(anyEvents), nil
	}
	// Strict-sync: caller wants delivery confirmation per call.
	req := NewBatchIngestRequest(events)
	_, resp, err := c.CostEventsAPIService.IngestEventsBatch(ctx).BatchIngestRequest(*req).Execute()
	if err != nil {
		if resp != nil && isTerminalStatusCode(resp.StatusCode) {
			return 0, err
		}
		return 0, fmt.Errorf("moolabs: cost ingest failed and buffer disabled: %w", err)
	}
	return len(events), nil
}

// IngestEvent is the ergonomic struct-arg cost-lane ingest method (US-010 —
// SDK Unified Ingest Methods PRD §3.3). Mirror of Python
// _CostNamespace.ingest_event and TypeScript CostNamespace.ingestEvent.
//
// Builds a CloudEvent envelope from args (no MeterSlug, no Value — cost
// lane carries data.spans[] instead) and routes to METER, not acute:
// unified ingest at the wire layer. Behavior is otherwise identical to
// UsageNamespace.IngestEvent:
//
//   - When the meter buffer is enabled (default), enqueues via the meter
//     producer channel (~50 ns) and returns IngestResult{Transport:
//     TransportBuffered}.
//   - When the meter buffer is disabled (Config.Buffer = &false), POSTs
//     inline via the meter F2 chain and returns IngestResult{Transport:
//     TransportSync}.
//
// Boundary checks (FR-6) fire synchronously BEFORE buffer enqueue:
//   - empty EventType / CustomerID / EntityID
//   - empty Spans slice (cost lane requires at least one span)
//   - missing / empty span_id on any span
//   - non-JSON-serializable Meta
//
// TenantID is NOT a field on IngestCostEventArgs (FR-3).
//
// NOTE: The legacy CostNamespace.IngestEvents batch method (above) still
// routes to acute and is retained for backward compatibility. This new
// IngestEvent method is the going-forward unified-surface entrypoint.
func (c *CostNamespace) IngestEvent(ctx context.Context, args IngestCostEventArgs) (IngestResult, error) {
	if len(args.Spans) == 0 {
		return IngestResult{}, errors.New("spans must contain at least one span")
	}
	envelope, err := buildEnvelope(costArgsToBuildEnvelopeArgs(args))
	if err != nil {
		return IngestResult{}, err
	}
	acceptedAt := time.Now().UTC()
	eventID, _ := envelope["id"].(string)

	if c.meterBuffer != nil {
		_ = c.meterProducer.SubmitAndCount([]any{envelope})
		return IngestResult{
			EventID:    eventID,
			Transport:  TransportBuffered,
			AcceptedAt: acceptedAt,
		}, nil
	}

	// Strict-sync: caller wants delivery confirmation per call.
	url := c.meterResolver.GetIngestURL()
	if _, postErr := postEventsBatch(ctx, url, c.apiKey, []any{envelope}); postErr != nil {
		if isTerminalIngestError(postErr) {
			return IngestResult{}, postErr
		}
		c.meterResolver.ReportPostOutcome(url, false)
		return IngestResult{}, fmt.Errorf("moolabs: cost ingest failed and buffer disabled: %w", postErr)
	}
	c.meterResolver.ReportPostOutcome(url, true)
	return IngestResult{
		EventID:    eventID,
		Transport:  TransportSync,
		AcceptedAt: acceptedAt,
	}, nil
}

// ── EventsNamespace (US-011 — dual-lane unified ingest) ──────────────────

// EventsNamespace is the standalone dual-lane ingest namespace. It does
// NOT embed any openapi-generated service — there's no "events" backing
// API in CapabilityMap. The single method client.Events.Ingest accepts
// IngestArgs (with optional MeterSlug + Value for the usage lane and
// optional Spans for the cost lane), runs the empty-lane guard, and
// routes a CloudEvent envelope through the same meter F2 + buffer chain
// as UsageNamespace.IngestEvent and CostNamespace.IngestEvent.
//
// Mirrors Python _EventsNamespace (US-004) and TypeScript EventsNamespace
// (US-008). Cross-language parity (US-013) asserts the same lane semantics
// across all three SDKs.
type EventsNamespace struct {
	// Meter-routing hooks. Identical to UsageNamespace's set — populated
	// by wireCapabilities at Moolabs construction.
	meterResolver *IngestURLResolver
	meterBuffer   *IngestBuffer
	meterProducer *ProducerChannel
	apiKey        string
}

// Ingest is the dual-lane ergonomic entrypoint (US-011).
//
// Empty-lane guard (FR §3.4): synchronous error when BOTH the usage lane
// is incomplete (MeterSlug == nil OR Value == nil) AND the cost lane is
// absent (len(Spans) == 0). Customers who want JUST usage pass
// MeterSlug + Value; customers who want JUST cost pass Spans; customers
// emitting both lanes for a single entity (e.g., an LLM call with both
// a counter delta AND per-span breakdown) pass all three.
//
// Boundary checks (FR-6) fire synchronously BEFORE buffer enqueue: empty
// EventType / CustomerID / EntityID, non-finite *Value, missing /
// empty / non-string span_id on any span, non-JSON-serializable Meta.
//
// Routes to the meter F2 chain (same target as Usage.IngestEvent and
// Cost.IngestEvent), so cost-shape envelopes and usage-shape envelopes
// hit a single unified ingest endpoint.
func (e *EventsNamespace) Ingest(ctx context.Context, args IngestArgs) (IngestResult, error) {
	// Empty-lane guard. Mirrors the Python _EventsNamespace.ingest check
	// and the TypeScript EventsNamespace.ingest check — the error message
	// is intentionally identical wording across the three SDKs (US-013
	// cross-language parity test will assert this).
	if !args.IsUsageLanePresent() && !args.IsCostLanePresent() {
		return IngestResult{}, errors.New("at least one lane required: pass MeterSlug + Value (usage), Spans (cost), or both")
	}
	envelope, err := buildEnvelope(ingestArgsToBuildEnvelopeArgs(args))
	if err != nil {
		return IngestResult{}, err
	}
	acceptedAt := time.Now().UTC()
	eventID, _ := envelope["id"].(string)

	if e.meterBuffer != nil {
		_ = e.meterProducer.SubmitAndCount([]any{envelope})
		return IngestResult{
			EventID:    eventID,
			Transport:  TransportBuffered,
			AcceptedAt: acceptedAt,
		}, nil
	}

	// Strict-sync: caller wants delivery confirmation per call.
	url := e.meterResolver.GetIngestURL()
	if _, postErr := postEventsBatch(ctx, url, e.apiKey, []any{envelope}); postErr != nil {
		if isTerminalIngestError(postErr) {
			return IngestResult{}, postErr
		}
		e.meterResolver.ReportPostOutcome(url, false)
		return IngestResult{}, fmt.Errorf("moolabs: events ingest failed and buffer disabled: %w", postErr)
	}
	e.meterResolver.ReportPostOutcome(url, true)
	return IngestResult{
		EventID:    eventID,
		Transport:  TransportSync,
		AcceptedAt: acceptedAt,
	}, nil
}

// stopProducer signals the cost producer to exit + waits for drain. Called
// from Moolabs.Close().
func (c *CostNamespace) stopProducer() {
	if c.producer == nil {
		return // never constructed
	}
	c.producer.Stop()
}

// IngestQueueDropped returns cost events dropped at the producer-channel layer
// (channel full at Submit). Distinct from buffer-layer drop_oldest.
func (c *CostNamespace) IngestQueueDropped() int64 {
	if c.producer == nil {
		return 0
	}
	return c.producer.IngestQueueDropped()
}

// ── internals ──

func (m *Moolabs) makeClientAtURL(host string) *APIClient {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: host}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+m.apiKey)
	// Deliberately do NOT install m.timeout as a client-wide HTTP timeout here.
	// m.timeout is the *ingest* drain budget; it's applied per-attempt via
	// context.WithTimeout in bufferDrain/costBufferDrain. Forcing it on the
	// shared generated clients would also cap NON-ingest calls (billing,
	// reports, credits, subscriptions, …) at the ingest ceiling, breaking
	// legitimately long customer operations. Non-ingest calls are bounded by
	// the caller's own context instead.
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
	m.Notifications = &NotificationsNamespace{
		NotificationsAPIService: m.meterClient.NotificationsAPI,
		AlertsAPIService:        m.bffClient.AlertsAPI,
	}
	// Usage capability first — the meter F2 resolver + buffer + producer
	// it owns are shared with Cost (US-010) for the unified IngestEvent
	// surface that posts cost-shape CloudEvent envelopes to meter (not
	// acute). The legacy CostEventsAPIService.IngestEventsBatch path
	// remains routed to acute via the separate costBuf below.
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

	// EventsNamespace (US-011) shares the SAME meter F2 chain + buffer +
	// producer as UsageNamespace. The dual-lane Ingest method routes a
	// single CloudEvent envelope through this shared pipeline; the
	// EventsNamespace doesn't own any additional resources.
	m.Events = &EventsNamespace{
		meterResolver: m.ingestResolver,
		meterBuffer:   buf,
		meterProducer: usage.producer,
		apiKey:        m.apiKey,
	}

	costBuf := m.lazyCostBuffer()
	cost := &CostNamespace{
		// sdk-cost-capability-acute-backing US-008 (Ralph iter 23 follow-up):
		// route client.cost.* DIRECT to acute.{baseURL} instead of the BFF
		// cost-ingest-proxy. Two backing services per HLD Appendix D.5 / RFC
		// + the parity-check'd dx_routing.go CapabilityMap["cost"] entry.
		CostEventsAPIService: m.acuteClient.CostEventsAPI,
		SdkIngestAPIService:  m.acuteClient.SdkIngestAPI,
		buffer:               costBuf,
		// US-010: Cost.IngestEvent routes to METER (not acute). Share the
		// usage resolver + buffer + producer + apiKey so both capabilities
		// hit the same unified ingest endpoint with shared backpressure.
		meterResolver: m.ingestResolver,
		meterBuffer:   buf,
		meterProducer: usage.producer,
		apiKey:        m.apiKey,
		logger:        m.logger,
	}
	// Producer-channel only wired when the buffer is enabled. Channel cap
	// scales with BufferMax: max(1024, BufferMax/8), same as usage.
	if costBuf != nil {
		chanCap := 1024
		if m.bufferMax/8 > chanCap {
			chanCap = m.bufferMax / 8
		}
		cost.producer = NewProducerChannel(chanCap, func(events []any) {
			_ = costBuf.Enqueue(events)
		}, m.logger)
	}
	m.Cost = cost
}

func (m *Moolabs) lazyBuffer() *IngestBuffer {
	if !m.bufferEnabled {
		return nil
	}
	m.bufferOnce.Do(func() {
		cfg := DefaultIngestBufferConfig
		cfg.MaxSize = m.bufferMax
		cfg.Logger = m.logger // propagate customer-supplied logger (NoopLogger if none)
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
	if m.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, m.timeout)
		defer cancel()
	}
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

// lazyCostBuffer constructs + starts the cost buffer on first use. Mirrors
// lazyBuffer; separate instance because cost drains to a different endpoint.
func (m *Moolabs) lazyCostBuffer() *IngestBuffer {
	if !m.bufferEnabled {
		return nil
	}
	m.costBufferOnce.Do(func() {
		cfg := DefaultIngestBufferConfig
		cfg.MaxSize = m.bufferMax
		cfg.Logger = m.logger // NoopLogger if customer provided none
		buf, err := NewIngestBuffer(m.costBufferDrain, cfg)
		if err != nil {
			panic("moolabs: building cost IngestBuffer: " + err.Error())
		}
		// Publish before Start (same ordering rationale as lazyBuffer: the
		// drain reads m.costBuffer for terminal-drop bumps).
		m.costBuffer = buf
		buf.Start()
	})
	return m.costBuffer
}

// costBufferDrain is the cost buffer's drain callback. It POSTs queued cost
// events to ACUTE's batch endpoint via the generated CostEventsAPI client.
//
// Terminal errors (4xx auth/validation) DISCARD the batch — retry fails
// identically; the count is recorded and a WARN logged (only surfaced if the
// customer supplied a Logger). Transient errors return 0 so the batch is
// retried next tick. Retry is idempotency-key safe (backend dedups).
func (m *Moolabs) costBufferDrain(ctx context.Context, events []any) (int, error) {
	if m.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, m.timeout)
		defer cancel()
	}
	costEvents := make([]CostEventIngest, 0, len(events))
	for _, e := range events {
		if ce, ok := e.(CostEventIngest); ok {
			costEvents = append(costEvents, ce)
		}
	}
	if len(costEvents) == 0 {
		// Nothing valid to send — treat as drained so a malformed entry can
		// never wedge the queue (defensive; should not happen via IngestEvents).
		return len(events), nil
	}
	req := NewBatchIngestRequest(costEvents)
	_, resp, err := m.acuteClient.CostEventsAPI.IngestEventsBatch(ctx).BatchIngestRequest(*req).Execute()
	if err != nil {
		status := 0
		if resp != nil {
			status = resp.StatusCode
		}
		if isTerminalStatusCode(status) {
			if m.costBuffer != nil {
				m.costBuffer.recordTerminalDrop(len(events))
			}
			m.logger.Warn("moolabs.cost_buffer.terminal_drop",
				"status", status,
				"count", len(events),
				"err", err,
			)
			return len(events), nil // discard — retry would fail identically
		}
		return 0, err // transient — re-enqueue + retry next tick
	}
	return len(events), nil
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
