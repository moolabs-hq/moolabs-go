// Ergonomic ingest envelope — leaf-pure helpers shared by US-009 / US-010 /
// US-011 of the SDK Unified Ingest Methods PRD:
// docs/prd/2026-06-02-sdk-unified-ingest-methods-prd.md.
//
// This file is a LEAF module — it depends only on the standard library, so
// it builds standalone under the CI leaf-test job in .github/workflows/
// sdk-pr.yml (no openapi-generated tree required). The three namespace
// methods that ATTACH this envelope to UsageNamespace / CostNamespace /
// Moolabs.Events live in dx_client.go because they reference the namespace
// types which embed openapi-generated services.
//
// Wire shape parity is enforced across Python (_dx_namespaces.py) and
// TypeScript (_dx_namespaces.ts) by the cross-language parity gate
// (US-013). entity_id maps to data.request_id — the threading key
// preserved for moo-meter's request_id column and acute's cross-lane join.
//
// TenantID is NOT a field on any *Args struct (FR-3). The server derives
// tenant identity from the API key.

package moolabs

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"
)

// defaultSDKSource is the source value stamped on envelopes when the
// caller does not pin their own app-level source. Matches the Python
// _DEFAULT_SDK_SOURCE constant and the TS DEFAULT_SDK_SOURCE constant.
const defaultSDKSource = "moolabs-sdk"

// IngestResult is the return type for the three new ergonomic ingest
// methods. Mirrors the Python @dataclass(frozen=True) IngestResult and the
// TypeScript IngestResult interface. Field names are PascalCase per Go
// convention; wire-side parity is enforced at the envelope layer, not here.
//
//   - EventID    — the id stamped on the envelope (auto-generated when the
//     caller did not pass IngestEventArgs.EventID). Useful for cross-lane
//     join with a sibling cost event sharing the same EntityID.
//   - Transport  — "buffered" when enqueued to the in-memory G5 buffer
//     (non-blocking; HTTP happens later via the drain worker), or "sync"
//     when posted on the caller's goroutine (strict-sync mode).
//   - AcceptedAt — client-side timestamp the SDK recorded when the envelope
//     was built / enqueued. NOT the server's receipt time.
type IngestResult struct {
	EventID    string
	Transport  string
	AcceptedAt time.Time
}

// Transport constants — values of IngestResult.Transport. Exported so
// callers can pattern-match without string literals.
const (
	TransportBuffered = "buffered"
	TransportSync     = "sync"
)

// IngestEventArgs is the struct-arg shape for UsageNamespace.IngestEvent
// (US-009). Required fields: EventType, CustomerID, EntityID, MeterSlug,
// Value. Optional fields: EventID, Source, Time, Meta.
//
// Note: Value is a float64 (not a pointer). The usage lane always carries
// a numeric value — zero is a legitimate counter delta and is therefore
// accepted; only NaN and +/-Inf are rejected by the boundary check.
//
// TenantID is intentionally absent (FR-3).
type IngestEventArgs struct {
	EventType  string
	CustomerID string
	EntityID   string
	MeterSlug  string
	Value      float64

	// Well-known top-level data.* keys (canonical wire shape).
	// Pointers so a zero/empty value means "explicitly zero" vs absent.
	Provider          *string
	Model             *string
	TotalInputTokens  *int64
	TotalOutputTokens *int64
	TotalTokens       *int64
	LatencyMs         *int64
	Status            *string

	EventID string
	Source  string
	Time    time.Time
	Meta    map[string]any
}

// IngestCostEventArgs is the struct-arg shape for CostNamespace.IngestEvent
// (US-010). Mirror of Python US-003's _CostNamespace.ingest_event kwargs and
// TypeScript US-007's CostIngestEventArgs. Required: EventType, CustomerID,
// EntityID, Spans (at least one span with non-empty span_id). Optional:
// EventID, Source, Time, Meta.
//
// No MeterSlug, no Value — those are usage-lane fields. The cost lane
// carries per-span breakdowns instead.
//
// TenantID is intentionally absent (FR-3).
type IngestCostEventArgs struct {
	EventType  string
	CustomerID string
	EntityID   string
	Spans      []map[string]any

	// Well-known top-level data.* keys (canonical wire shape).
	Provider          *string
	Model             *string
	TotalInputTokens  *int64
	TotalOutputTokens *int64
	TotalTokens       *int64
	LatencyMs         *int64
	Status            *string

	EventID string
	Source  string
	Time    time.Time
	Meta    map[string]any
}

// costArgsToBuildEnvelopeArgs translates the public CostNamespace.IngestEvent
// args to the internal buildEnvelope shape. Extracted as a leaf-pure helper
// so the mapping itself is testable without constructing a CostNamespace
// (which embeds openapi-generated services and therefore can't be built in
// the leaf-test compile unit).
func costArgsToBuildEnvelopeArgs(args IngestCostEventArgs) buildEnvelopeArgs {
	return buildEnvelopeArgs{
		EventType:  args.EventType,
		CustomerID: args.CustomerID,
		EntityID:   args.EntityID,
		// MeterSlug deliberately empty: cost lane has no meter routing key.
		// HasValue deliberately false: cost lane has no scalar value.
		Spans: args.Spans,
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
	}
}

// IngestArgs is the struct-arg shape for EventsNamespace.Ingest (US-011 —
// dual-lane unified entrypoint). Mirror of Python US-004
// _EventsNamespace.ingest kwargs and TypeScript US-008 EventsIngestArgs.
//
// Required: EventType, CustomerID, EntityID.
// Usage lane (both required for that lane to be present): *MeterSlug, *Value.
// Cost lane (present iff non-empty): Spans.
// Optional: EventID, Source, Time, Meta.
//
// Pointers are used for the usage-lane optionals so the caller can
// distinguish "not provided" from "zero value." A *float64 of nil means
// "no value supplied"; a *float64 pointing at 0 is a legitimate zero
// counter delta. Same logic for *MeterSlug: an empty-string pointer would
// be ambiguous; nil = "not provided."
//
// The empty-lane guard (FR §3.4) fires when BOTH the usage lane is
// incomplete AND the cost lane is absent — i.e., the caller supplied
// neither (meter_slug + value) nor any spans.
//
// TenantID is intentionally absent (FR-3).
type IngestArgs struct {
	EventType  string
	CustomerID string
	EntityID   string

	MeterSlug *string
	Value     *float64
	Spans     []map[string]any

	// Well-known top-level data.* keys (canonical wire shape).
	Provider          *string
	Model             *string
	TotalInputTokens  *int64
	TotalOutputTokens *int64
	TotalTokens       *int64
	LatencyMs         *int64
	Status            *string

	EventID string
	Source  string
	Time    time.Time
	Meta    map[string]any
}

// IsUsageLanePresent returns true when the caller supplied both MeterSlug
// and Value (i.e., the usage-lane required fields). Pure helper exposed
// so callers + the empty-lane guard share one definition of "lane present."
func (a IngestArgs) IsUsageLanePresent() bool {
	return a.MeterSlug != nil && a.Value != nil
}

// IsCostLanePresent returns true when the caller supplied at least one
// span. An empty (or nil) Spans slice means the cost lane is absent.
func (a IngestArgs) IsCostLanePresent() bool {
	return len(a.Spans) > 0
}

// ingestArgsToBuildEnvelopeArgs translates the public EventsNamespace.Ingest
// args to the internal buildEnvelope shape. Extracted as a leaf-pure
// helper so the lane-handling mapping is testable without constructing an
// EventsNamespace.
//
// The empty-lane guard MUST fire upstream (in EventsNamespace.Ingest)
// BEFORE this helper runs — this helper assumes the args have at least
// one lane present and produces an envelope honoring whatever lanes were
// supplied. If neither lane is present this helper still returns a
// usable buildEnvelopeArgs (with no Value, no MeterSlug, no Spans) and
// buildEnvelope would accept it; the resulting envelope is meaningless
// data-wise, so the upstream guard exists to fail fast at the call site.
func ingestArgsToBuildEnvelopeArgs(args IngestArgs) buildEnvelopeArgs {
	be := buildEnvelopeArgs{
		EventType:  args.EventType,
		CustomerID: args.CustomerID,
		EntityID:   args.EntityID,
		Spans:      args.Spans,
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
	}
	if args.MeterSlug != nil {
		be.MeterSlug = *args.MeterSlug
	}
	if args.Value != nil {
		be.HasValue = true
		be.Value = *args.Value
	}
	return be
}

// newEventID returns a 32-char lowercase hex id (uuid4-equivalent — 128
// random bits, no dashes). Matches the Python uuid4().hex format and the
// TS dash-stripped randomUUID format so the wire-side id field is
// grep-compatible across all three language SDKs.
func newEventID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", fmt.Errorf("moolabs: generating event_id: %w", err)
	}
	return hex.EncodeToString(b[:]), nil
}

// checkValueIsFinite rejects NaN and +/-Inf for the usage-lane Value field
// (FR-6). Matches the Python _check_value_is_finite contract; the Go
// signature drops the bool/non-numeric branch because float64 cannot
// represent those at compile time.
func checkValueIsFinite(value float64, name string) error {
	if math.IsNaN(value) {
		return fmt.Errorf("%s must be a finite number; got NaN", name)
	}
	if math.IsInf(value, 0) {
		return fmt.Errorf("%s must be a finite number; got %v", name, value)
	}
	return nil
}

// checkSpansHaveSpanIDs ensures every span carries a non-empty span_id.
// Acute's per-span dedup grain is "sdk:{span_id}" — an empty/missing
// span_id collides every cost span into the same dedup key, silently
// dropping legitimate spans. Surface up front so the bug is visible at
// the call site rather than as a downstream "where did my cost data go"
// investigation. Matches Python _check_spans_have_span_ids.
func checkSpansHaveSpanIDs(spans []map[string]any) error {
	for i, span := range spans {
		if span == nil {
			return fmt.Errorf("spans[%d] must be a non-nil map", i)
		}
		raw, present := span["span_id"]
		if !present {
			return fmt.Errorf("spans[%d].span_id must be a non-empty string", i)
		}
		spanID, ok := raw.(string)
		if !ok || spanID == "" {
			return fmt.Errorf("spans[%d].span_id must be a non-empty string", i)
		}
	}
	return nil
}

// checkCostSpansHaveProviderAndModel ensures every cost-lane span carries
// a non-empty provider and model. Acute's cost-enricher rejects spans
// without both via a silent continue at
// services/moo-acute/app/workers/cost_enricher.py; events with all-
// invalid spans produce NO log line at acute, making the drop invisible
// to operators. Surface up front so the failure surfaces at the call site
// rather than as a downstream "cost event silently disappeared"
// investigation. Matches Python _check_cost_spans_have_provider_and_model.
func checkCostSpansHaveProviderAndModel(spans []map[string]any) error {
	for i, span := range spans {
		if span == nil {
			// Caught by checkSpansHaveSpanIDs already; defensive only.
			continue
		}
		provider, ok := span["provider"].(string)
		if !ok || provider == "" {
			return fmt.Errorf(
				"spans[%d].provider must be a non-empty string — spans without "+
					"provider are silently dropped during downstream cost "+
					"processing, resulting in missing cost-attribution data",
				i,
			)
		}
		model, ok := span["model"].(string)
		if !ok || model == "" {
			return fmt.Errorf(
				"spans[%d].model must be a non-empty string — spans without "+
					"model are silently dropped during downstream cost "+
					"processing, resulting in missing cost-attribution data",
				i,
			)
		}
	}
	return nil
}

// checkMetaIsJSONSerializable verifies that meta round-trips through JSON
// BEFORE buffer enqueue. If meta carries a non-serializable value (e.g. a
// channel, a func, or a map keyed by a non-string), the failure would
// otherwise happen in the drain worker AFTER the customer's call already
// returned — an async failure mode the customer cannot catch. Surface it
// synchronously here. Matches Python _check_meta_is_json_serializable.
func checkMetaIsJSONSerializable(meta map[string]any) error {
	if _, err := json.Marshal(meta); err != nil {
		return fmt.Errorf("meta must be JSON-serializable: %w", err)
	}
	return nil
}

// buildEnvelopeArgs is the internal envelope-builder input. The internal
// HasValue flag lets the cost-lane (US-010) and dual-lane (US-011) callers
// build envelopes without a Value field, while keeping the usage-lane
// (US-009) caller's required-value semantics. Exposed only inside this
// package.
type buildEnvelopeArgs struct {
	EventType  string
	CustomerID string
	EntityID   string

	MeterSlug string
	HasValue  bool
	Value     float64
	Spans     []map[string]any

	// Well-known top-level data.* keys (canonical wire shape). Pointer-
	// typed so callers can pass nil to OMIT the key from data, vs
	// supplying a non-nil pointer (even to zero/empty) to include it.
	Provider          *string
	Model             *string
	TotalInputTokens  *int64
	TotalOutputTokens *int64
	TotalTokens       *int64
	LatencyMs         *int64
	Status            *string

	EventID string
	Source  string
	Time    time.Time
	Meta    map[string]any
}

// buildEnvelope assembles a CloudEvents 1.0 envelope from ergonomic args.
//
// Wire contract: docs/prd/2026-06-02-sdk-unified-ingest-methods-prd.md
// Section 4 canonical envelope. EntityID maps to data.request_id on the
// wire (the threading key preserved for moo-meter's request_id column and
// acute's cross-lane join).
//
// Boundary checks (FR-6) fire synchronously BEFORE buffer enqueue, so
// customers see the rejection at their call site:
//
//   - empty EventType / CustomerID / EntityID
//   - non-finite Value (NaN / +Inf / -Inf)
//   - missing or empty span_id on any span
//   - non-JSON-serializable Meta
//
// Returns a map[string]any (the wire shape) directly rather than a
// generated Event struct — the F2 fallback chain marshals []any with
// json.Marshal anyway, so a generated-struct round-trip would add zero
// type safety and one allocation per event.
func buildEnvelope(args buildEnvelopeArgs) (map[string]any, error) {
	if args.EventType == "" {
		return nil, errors.New("event_type must be a non-empty string")
	}
	if args.CustomerID == "" {
		return nil, errors.New("customer_id must be a non-empty string")
	}
	if args.EntityID == "" {
		return nil, errors.New("entity_id must be a non-empty string")
	}
	if args.HasValue {
		if err := checkValueIsFinite(args.Value, "value"); err != nil {
			return nil, err
		}
	}
	if args.Spans != nil {
		if err := checkSpansHaveSpanIDs(args.Spans); err != nil {
			return nil, err
		}
		// Lane discriminator: USAGE envelopes carry MeterSlug+Value (HasValue=
		// true) and may optionally include Spans as supplemental context (e.g.
		// event lineage or per-span breakdown attached to a usage emit). Those
		// spans don't flow through the downstream cost-enrichment pipeline and
		// don't need provider/model. COST envelopes carry ONLY Spans (no
		// MeterSlug, no HasValue) and DO flow to the cost-enrichment pipeline,
		// which silently drops spans without provider+model. The provider+
		// model check fires only on the cost lane.
		isCostLane := args.MeterSlug == "" && !args.HasValue
		if isCostLane {
			topLevelProvider := args.Provider != nil && *args.Provider != ""
			topLevelModel := args.Model != nil && *args.Model != ""
			if !topLevelProvider || !topLevelModel {
				if err := checkCostSpansHaveProviderAndModel(args.Spans); err != nil {
					return nil, err
				}
			}
		}
	}
	if args.Meta != nil {
		if err := checkMetaIsJSONSerializable(args.Meta); err != nil {
			return nil, err
		}
	}

	// Assemble data. Well-known top-level keys (provider, model,
	// total_*_tokens, latency_ms, status) land at data.<key> directly per
	// the canonical wire shape. Arbitrary customer fields go through
	// args.Meta and land nested at data.meta.<key> per Decision 4 of the
	// PRD's HOW section.
	data := map[string]any{"request_id": args.EntityID}
	if args.MeterSlug != "" {
		data["meter_slug"] = args.MeterSlug
	}
	if args.HasValue {
		data["value"] = args.Value
	}
	// Well-known top-level data.* keys (canonical wire shape).
	if args.Provider != nil {
		data["provider"] = *args.Provider
	}
	if args.Model != nil {
		data["model"] = *args.Model
	}
	if args.TotalInputTokens != nil {
		data["total_input_tokens"] = *args.TotalInputTokens
	}
	if args.TotalOutputTokens != nil {
		data["total_output_tokens"] = *args.TotalOutputTokens
	}
	if args.TotalTokens != nil {
		data["total_tokens"] = *args.TotalTokens
	}
	if args.LatencyMs != nil {
		data["latency_ms"] = *args.LatencyMs
	}
	if args.Status != nil {
		data["status"] = *args.Status
	}
	if args.Spans != nil {
		data["spans"] = args.Spans
	}
	if args.Meta != nil {
		data["meta"] = args.Meta
	}

	eventID := args.EventID
	if eventID == "" {
		generated, err := newEventID()
		if err != nil {
			return nil, err
		}
		eventID = generated
	}
	source := args.Source
	if source == "" {
		source = defaultSDKSource
	}
	t := args.Time
	if t.IsZero() {
		t = time.Now().UTC()
	}

	return map[string]any{
		"id":              eventID,
		"specversion":     "1.0",
		"source":          source,
		"type":            args.EventType,
		"subject":         args.CustomerID,
		"time":            t.UTC().Format(time.RFC3339Nano),
		"datacontenttype": "application/json",
		"data":            data,
	}, nil
}

// UsageNamespace.IngestEvent (US-009) is defined alongside the existing
// IngestEvents in dx_client.go because it references the openapi-generated
// types embedded in UsageNamespace.
