// Capability → backing-service routing map for the normalized Moolabs Go SDK.
//
// Go port of sdks/dx/python/moolabs/_dx_routing.py — cross-language parity
// asserted by the automated parity job (BE Task H). Drift between Python,
// TypeScript, and Go here fails the build.
//
// See _dx_routing.py for full design notes; this file is intentionally a
// one-to-one mirror.

package moolabs

import "fmt"

// Backend identifies which Moolabs service hosts a backing API.
type Backend string

const (
	BackendBFF   Backend = "bff"
	BackendMeter Backend = "meter"
	BackendArc   Backend = "arc"
	// BackendAcute — sdk-cost-capability-acute-backing US-008 (T-11).
	// Drives client.cost.* direct to acute.{baseURL} instead of through
	// the BFF cost-ingest-proxy router.
	BackendAcute Backend = "acute"
)

// BackingClass is one openapi-generator-emitted service that backs a
// capability namespace. The APIServiceField is the name of the field
// on *APIClient that holds the corresponding *<X>APIService — the Go
// generator emits these as struct fields (e.g. APIClient.WalletsAPI).
type BackingClass struct {
	APIServiceField string // e.g. "WalletsAPI" (matches APIClient.WalletsAPI field)
	Backend         Backend
}

// SubdomainMap maps a backend to its DNS subdomain prefix. Customer's
// baseURL is the root domain (default "moolabs.com"); the SDK prepends
// "https://<subdomain>." per call. Customer code never types any of these.
var SubdomainMap = map[Backend]string{
	BackendBFF:   "api",
	BackendMeter: "meter",
	BackendArc:   "arc",
	BackendAcute: "acute",
}

// RegionIngestMap matches BFF's REGION_INGEST_MAP in
// services/moolabs-app/bff/app/api/v1/tenant_config.py. F2 fallback chain
// step 3 composes "https://ingest.{regionCode}.{baseURL}" from this map
// plus a region source (today: DefaultRegion; future: extracted from API key).
var RegionIngestMap = map[string]string{
	"us-east-1":      "us",
	"us-west-2":      "us",
	"ap-southeast-1": "ap",
	"eu-west-1":      "eu",
}

// DefaultRegion is the fallback when the SDK has no other region signal.
const DefaultRegion = "us-east-1"

// CapabilityMap defines the 11 customer-facing capability namespaces and the
// generated API-service struct fields that back each one. The Go SDK uses
// struct embedding to promote methods from these services onto the
// capability namespace structs in dx_client.go.
//
// O4 (verified 2026-05-15 against moolabs-py@main) found zero method-name
// collisions across all multi-class capabilities — so struct embedding works
// without ambiguity. Adding an Arc backing class whose method names collide
// with an existing one will surface as a Go compile error on the embedded
// struct; that's the right time to apply the sub-accessor pattern.
var CapabilityMap = map[string][]BackingClass{
	"usage": {
		{APIServiceField: "EventsAPI", Backend: BackendMeter},
		{APIServiceField: "MetersAPI", Backend: BackendMeter},
	},
	"customers": {
		{APIServiceField: "CustomersAPI", Backend: BackendMeter},
		{APIServiceField: "SubjectsAPI", Backend: BackendMeter},
	},
	"catalog": {
		{APIServiceField: "ProductCatalogAPI", Backend: BackendMeter},
		{APIServiceField: "RateCardsAPI", Backend: BackendBFF},
	},
	"subscriptions": {
		{APIServiceField: "MeterSubscriptionsAPI", Backend: BackendMeter},
	},
	"entitlements": {
		{APIServiceField: "EntitlementsAPI", Backend: BackendMeter},
	},
	"wallets": {
		{APIServiceField: "WalletsAPI", Backend: BackendBFF},
	},
	"credits": {
		{APIServiceField: "GrantsAPI", Backend: BackendBFF},
		{APIServiceField: "LedgerAPI", Backend: BackendBFF},
		{APIServiceField: "AutoTopupAPI", Backend: BackendBFF},
	},
	"billing": {
		{APIServiceField: "MeterBillingAPI", Backend: BackendMeter},
		{APIServiceField: "RatingAPI", Backend: BackendBFF},
		{APIServiceField: "FxRatesAPI", Backend: BackendBFF},
	},
	"collections": {
		{APIServiceField: "AccountsAPI", Backend: BackendArc},
		{APIServiceField: "AccountTeamAPI", Backend: BackendArc},
		{APIServiceField: "AnalyticsAPI", Backend: BackendArc},
		{APIServiceField: "CasesAPI", Backend: BackendArc},
		{APIServiceField: "CashCreditsAPI", Backend: BackendArc},
		{APIServiceField: "ArcCommunicationsAPI", Backend: BackendArc}, // stitcher-renamed tag
		{APIServiceField: "CreditMemosAPI", Backend: BackendArc},
		{APIServiceField: "DisputesAPI", Backend: BackendArc},
		{APIServiceField: "EscalationsAPI", Backend: BackendArc},
		{APIServiceField: "HandoffsAPI", Backend: BackendArc},
		{APIServiceField: "NotesAPI", Backend: BackendArc},
		{APIServiceField: "PaymentsAPI", Backend: BackendArc},
		{APIServiceField: "PlansAPI", Backend: BackendArc},
		{APIServiceField: "PromisesAPI", Backend: BackendArc},
		{APIServiceField: "RemittancesAPI", Backend: BackendArc},
		{APIServiceField: "ReportsAPI", Backend: BackendArc},
		{APIServiceField: "TasksAPI", Backend: BackendArc},
	},
	// AI cost-intelligence ingestion — DIRECT to acute.{baseURL}.
	// Per sdk-cost-capability-acute-backing US-008 (T-11). See
	// _dx_routing.py for full design rationale.
	//
	// CostEventsAPI backs:
	//   client.Cost.IngestEvent       (POST /api/v1/cost/ingest)
	//   client.Cost.IngestBatch       (POST /api/v1/cost/ingest/batch)
	//   client.Cost.SubmitAdjustment  (POST /api/v1/cost/adjustments)
	//
	// SdkIngestAPI backs:
	//   client.Cost.IngestSdkSpans    (POST /api/v1/ingest/batch)
	"cost": {
		{APIServiceField: "CostEventsAPI", Backend: BackendAcute},
		{APIServiceField: "SdkIngestAPI",  Backend: BackendAcute},
	},
	"notifications": {
		{APIServiceField: "NotificationsAPI", Backend: BackendMeter},
		{APIServiceField: "AlertsAPI", Backend: BackendBFF},
	},
}

// CapabilityOrder is the ordered list of capabilities for stable iteration
// (e.g. cross-language parity diffs). Same order as in contracts §3.2.
var CapabilityOrder = []string{
	"usage",
	"customers",
	"catalog",
	"subscriptions",
	"entitlements",
	"wallets",
	"credits",
	"billing",
	"collections",
	"cost",
	"notifications",
}

// init runs at package load. Mirrors _validate_routing() in the Python module
// (and validateRouting() in TS): fail fast on backend / order / empty-capability
// drift so a config bug surfaces at import, not on a mysterious method call.
func init() {
	if err := validateRouting(); err != nil {
		panic("moolabs: routing-map invariant violated: " + err.Error())
	}
}

func validateRouting() error {
	validBackends := make(map[Backend]struct{}, len(SubdomainMap))
	for b := range SubdomainMap {
		validBackends[b] = struct{}{}
	}
	capSet := make(map[string]struct{}, len(CapabilityMap))
	for c := range CapabilityMap {
		capSet[c] = struct{}{}
	}
	orderSet := make(map[string]struct{}, len(CapabilityOrder))
	for _, c := range CapabilityOrder {
		orderSet[c] = struct{}{}
	}
	if len(capSet) != len(orderSet) {
		return fmt.Errorf("CapabilityMap (%d) vs CapabilityOrder (%d) size drift",
			len(capSet), len(orderSet))
	}
	for c := range capSet {
		if _, ok := orderSet[c]; !ok {
			return fmt.Errorf("capability %q in CapabilityMap but not CapabilityOrder", c)
		}
	}
	for c := range orderSet {
		if _, ok := capSet[c]; !ok {
			return fmt.Errorf("capability %q in CapabilityOrder but not CapabilityMap", c)
		}
	}
	for capability, classes := range CapabilityMap {
		if len(classes) == 0 {
			return fmt.Errorf(
				"capability %q has no backing classes — every namespace must back at least one class",
				capability)
		}
		for _, bc := range classes {
			if _, ok := validBackends[bc.Backend]; !ok {
				return fmt.Errorf(
					"capability %q service %q declares backend %q not in SubdomainMap",
					capability, bc.APIServiceField, bc.Backend)
			}
		}
	}
	return nil
}
