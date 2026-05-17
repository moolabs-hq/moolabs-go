// URL derivation + F2 ingest fallback chain — Go port of _dx_urls.py.
//
// Two responsibilities:
//
//  1. DeriveHost(backend, baseURL) — convention-based subdomain rewriting.
//     Pure function, no state.
//
//  2. IngestURLResolver — stateful F2 fallback chain for the event-ingest
//     hot path (contracts §3.5). Resolves the URL to POST events to via a
//     4-step chain.
//
// Thread-safe: all state mutations are guarded by an internal sync.Mutex —
// Go programs typically have multiple goroutines posting events; this matters.

package moolabs

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	meterIngestPath = "/api/v1/events"
	// DiscoveryPath is the BFF endpoint the F2 resolver consults for the
	// regional ingest URL. Exported for the client to use directly.
	DiscoveryPath = "/v1/tenant/config"
)

// NormalizeBaseURL strips scheme/path from a caller-provided baseURL and
// returns the bare host. Accepts "moolabs.com", "https://moolabs.com",
// "https://moolabs.com/", or "  moolabs.com  ".
//
// Returns an error on empty/invalid input — typically wired so a customer
// typo crashes at construction.
func NormalizeBaseURL(baseURL string) (string, error) {
	stripped := strings.TrimRight(strings.TrimSpace(baseURL), "/")
	if stripped == "" {
		return "", fmt.Errorf("moolabs: baseURL is empty after stripping, got %q", baseURL)
	}
	withScheme := stripped
	if !strings.Contains(stripped, "://") {
		withScheme = "https://" + stripped
	}
	u, err := url.Parse(withScheme)
	if err != nil || u.Host == "" {
		return "", fmt.Errorf("moolabs: baseURL has no parseable host: %q", baseURL)
	}
	return u.Host, nil
}

// DeriveHost returns "https://<subdomain>.<baseURL>" for a backend.
func DeriveHost(backend Backend, baseURL string) (string, error) {
	subdomain, ok := SubdomainMap[backend]
	if !ok {
		return "", fmt.Errorf("moolabs: unknown backend %q; known: %v",
			backend, sortedSubdomains())
	}
	host, err := NormalizeBaseURL(baseURL)
	if err != nil {
		return "", err
	}
	return "https://" + subdomain + "." + host, nil
}

func sortedSubdomains() []string {
	out := make([]string, 0, len(SubdomainMap))
	for b := range SubdomainMap {
		out = append(out, string(b))
	}
	return out
}

// ── F2 ingest fallback chain ─────────────────────────────────────────────

// IngestResolverConfig tunes the F2 chain's behavior. Defaults match the
// Python and TS versions (contracts §3.5 / O6).
type IngestResolverConfig struct {
	// DiscoveryRetryTTL — after a failed discovery attempt, the SDK skips
	// re-trying discovery for this window.
	DiscoveryRetryTTL time.Duration

	// PostFailureThreshold — consecutive POST failures to a cached URL
	// before cache invalidation.
	PostFailureThreshold int

	// RecentlyFailedTTL — how long the SDK remembers a URL as "recently
	// failed" so a subsequent discovery returning it falls through to step 3.
	RecentlyFailedTTL time.Duration
}

// DefaultIngestResolverConfig holds the illustrative defaults pinned at LLD (O6).
var DefaultIngestResolverConfig = IngestResolverConfig{
	DiscoveryRetryTTL:    60 * time.Second,
	PostFailureThreshold: 3,
	RecentlyFailedTTL:    300 * time.Second, // 5 min
}

// DiscoveryFn returns the parsed /tenant/config response. The caller wraps
// the HTTP call; this module has no httplib dependency.
type DiscoveryFn func() (map[string]interface{}, error)

// Clock returns the current time. Tests inject a fake clock for determinism.
type Clock func() time.Time

// IngestURLResolver implements the F2 fallback chain for event ingest.
// Safe for concurrent use.
type IngestURLResolver struct {
	baseURL     string
	discoveryFn DiscoveryFn
	region      string
	config      IngestResolverConfig
	clock       Clock

	mu                    sync.Mutex
	cv                    *sync.Cond // singleflight discovery wait — uses &mu
	discoveryInFlight     bool       // true while one goroutine runs discoveryFn
	cachedURL             string
	discoveryBlockedUntil time.Time
	recentlyFailed        map[string]time.Time // URL → expiry
	postFailures          map[string]int       // URL → consecutive count
}

// NewIngestURLResolver constructs a resolver.
func NewIngestURLResolver(baseURL string, discoveryFn DiscoveryFn) (*IngestURLResolver, error) {
	host, err := NormalizeBaseURL(baseURL)
	if err != nil {
		return nil, err
	}
	r := &IngestURLResolver{
		baseURL:        host,
		discoveryFn:    discoveryFn,
		region:         DefaultRegion,
		config:         DefaultIngestResolverConfig,
		clock:          time.Now,
		recentlyFailed: make(map[string]time.Time),
		postFailures:   make(map[string]int),
	}
	r.cv = sync.NewCond(&r.mu)
	return r, nil
}

// WithRegion overrides the region used by step-3 fallback.
func (r *IngestURLResolver) WithRegion(region string) *IngestURLResolver {
	r.region = region
	return r
}

// WithConfig overrides the tunables. Useful for tests.
func (r *IngestURLResolver) WithConfig(cfg IngestResolverConfig) *IngestURLResolver {
	r.config = cfg
	return r
}

// WithClock overrides the clock. Useful for tests.
func (r *IngestURLResolver) WithClock(c Clock) *IngestURLResolver {
	r.clock = c
	return r
}

// GetIngestURL runs the F2 chain and returns a URL to POST events to.
// Always returns a URL; discovery failures fall through to step 3/4 rather
// than returning an error.
//
// Singleflight discovery (post-PR #395 review I3): the 10s HTTP discovery
// call is made WITHOUT holding the resolver mutex so concurrent operations
// don't serialize behind it. Other goroutines arriving while discovery is
// in flight wait on a sync.Cond for the discoverer to commit its result,
// then re-check the cache. A burst of N concurrent ingest goroutines
// triggers AT MOST one discovery.
func (r *IngestURLResolver) GetIngestURL() string {
	r.mu.Lock()
	r.expireRecentlyFailedLocked()

	// Step 1 — cache hit
	if r.cachedURL != "" {
		out := r.cachedURL
		r.mu.Unlock()
		return out
	}

	// Decide whether to discover.
	now := r.clock()
	shouldDiscover := r.discoveryFn != nil && !now.Before(r.discoveryBlockedUntil)
	if !shouldDiscover {
		out := r.regionFallbackURLLocked()
		r.mu.Unlock()
		return out
	}

	if r.discoveryInFlight {
		// Wait for the in-flight discoverer to commit. cv.Wait releases
		// the mutex and re-acquires on wake. The discoverer notifies all
		// after commit; each waiter re-checks the cache below.
		r.cv.Wait()
		var out string
		if r.cachedURL != "" {
			out = r.cachedURL
		} else {
			out = r.regionFallbackURLLocked()
		}
		r.mu.Unlock()
		return out
	}

	// We own the discovery. Mark in-flight, release the mutex for the HTTP.
	r.discoveryInFlight = true
	r.mu.Unlock()

	// Critical: the in-flight flag MUST be cleared and the cv MUST be
	// broadcast even on panic. Without this, a panic in the user-supplied
	// DiscoveryFn (which we don't control) leaves discoveryInFlight = true
	// forever; every subsequent GetIngestURL call blocks on cv.Wait()
	// expecting a Broadcast that will never come — the whole SDK hangs
	// (post-round-2 review C-NEW-1). The defer also runs on the normal
	// completion path; we set `discovered` first so the post-defer commit
	// logic uses it.
	var discovered string
	commit := func() {
		r.mu.Lock()
		r.discoveryInFlight = false
		if discovered != "" {
			r.cachedURL = discovered
		}
		r.cv.Broadcast()
		r.mu.Unlock()
	}
	defer func() {
		if rec := recover(); rec != nil {
			commit()
			panic(rec) // re-raise so the caller still sees the bug
		}
	}()

	discovered = r.tryDiscoveryUnlocked()
	commit()

	r.mu.Lock()
	var out string
	if r.cachedURL != "" {
		out = r.cachedURL
	} else {
		out = r.regionFallbackURLLocked()
	}
	r.mu.Unlock()
	return out
}

// ReportPostOutcome updates state based on a POST attempt's result.
//
// On success: reset per-URL failure counter.
// On failure: increment; at threshold AND if url is cached, invalidate
// cache AND record url as recently failed.
func (r *IngestURLResolver) ReportPostOutcome(postURL string, success bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if success {
		delete(r.postFailures, postURL)
		return
	}
	count := r.postFailures[postURL] + 1
	r.postFailures[postURL] = count
	if count >= r.config.PostFailureThreshold {
		if r.cachedURL == postURL {
			r.cachedURL = ""
		}
		r.recentlyFailed[postURL] = r.clock().Add(r.config.RecentlyFailedTTL)
		delete(r.postFailures, postURL)
		// Bound the maps (post-PR #395 review M2). Defends against
		// buggy BFF returning many distinct dead URLs faster than the
		// TTL expires them.
		r.capRecentlyFailedLocked(64)
		r.capPostFailuresLocked(64)
	}
}

// capRecentlyFailedLocked evicts entries with the earliest expiry when over
// budget. Caller must hold r.mu.
func (r *IngestURLResolver) capRecentlyFailedLocked(maxEntries int) {
	if len(r.recentlyFailed) <= maxEntries {
		return
	}
	// Collect URL+expiry pairs and sort ascending by expiry; drop oldest.
	type pair struct {
		url    string
		expiry time.Time
	}
	pairs := make([]pair, 0, len(r.recentlyFailed))
	for u, t := range r.recentlyFailed {
		pairs = append(pairs, pair{u, t})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].expiry.Before(pairs[j].expiry) })
	dropCount := len(r.recentlyFailed) - maxEntries
	for _, p := range pairs[:dropCount] {
		delete(r.recentlyFailed, p.url)
	}
}

// capPostFailuresLocked evicts the highest-count entries when over budget
// (they're closest to natural eviction at the failure threshold anyway).
// Caller must hold r.mu.
func (r *IngestURLResolver) capPostFailuresLocked(maxEntries int) {
	if len(r.postFailures) <= maxEntries {
		return
	}
	type pair struct {
		url   string
		count int
	}
	pairs := make([]pair, 0, len(r.postFailures))
	for u, c := range r.postFailures {
		pairs = append(pairs, pair{u, c})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].count > pairs[j].count })
	dropCount := len(r.postFailures) - maxEntries
	for _, p := range pairs[:dropCount] {
		delete(r.postFailures, p.url)
	}
}

// Step4LastResortURL returns "meter.{baseURL}/api/v1/events" — the always-
// derivable last-resort URL. Independent of region, discovery, or cache.
// Exported for tests + observability.
//
// DeriveHost can in principle return an error if BackendMeter is missing
// from subdomainMap, but that's impossible at runtime: BackendMeter is a
// compile-time constant and subdomainMap is initialized at package load.
// We panic on the "impossible" case rather than silently returning a
// malformed URL — production caller would otherwise post to "/api/v1/events"
// (no host) and quietly fail F2-step-4 forever.
func (r *IngestURLResolver) Step4LastResortURL() string {
	host, err := DeriveHost(BackendMeter, r.baseURL)
	if err != nil {
		// Programming error: subdomainMap regressed to exclude "meter".
		// Panic loudly — silent malformed URL is worse than crash.
		panic("moolabs: BUG: DeriveHost(BackendMeter, baseURL) failed: " + err.Error())
	}
	return host + meterIngestPath
}

// CachedURL returns the currently-cached ingest URL or "" if none.
func (r *IngestURLResolver) CachedURL() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.cachedURL
}

// ── internals ──

// tryDiscoveryUnlocked calls discoveryFn (HTTP, up to 10s) WITHOUT holding
// r.mu and returns the resolved URL or "" on any failure. State mutations
// that depend on the mutex (setting discoveryBlockedUntil on failure,
// checking recentlyFailed) are briefly re-acquired at the end.
//
// Post-PR #395 review I3: previous tryDiscoveryLocked held the mutex
// through the HTTP call, blocking every concurrent operation on the
// resolver for the full HTTP latency. The caller (GetIngestURL) now
// coordinates singleflight via cv so this function doesn't need to
// worry about duplicate discoveries.
func (r *IngestURLResolver) tryDiscoveryUnlocked() string {
	if r.discoveryFn == nil {
		return ""
	}
	blockAfter := false
	defer func() {
		if blockAfter {
			r.mu.Lock()
			r.discoveryBlockedUntil = r.clock().Add(r.config.DiscoveryRetryTTL)
			r.mu.Unlock()
		}
	}()

	response, err := r.discoveryFn()
	if err != nil {
		blockAfter = true
		return ""
	}
	endpoints, ok := response["endpoints"].(map[string]interface{})
	if !ok {
		blockAfter = true
		return ""
	}
	ingestHost, ok := endpoints["ingest"].(string)
	if !ok || ingestHost == "" {
		blockAfter = true
		return ""
	}
	// Normalize: BFF returns host-only ingest URL; append path if so.
	ingestHost = strings.TrimRight(ingestHost, "/")
	var fullURL string
	if strings.Contains(ingestHost, "://") &&
		!strings.Contains(strings.SplitN(ingestHost, "://", 2)[1], "/") {
		fullURL = ingestHost + meterIngestPath
	} else {
		fullURL = ingestHost
	}
	// Defense-in-depth: discovered URL must be under the customer's
	// configured base_url.
	if !hostMatchesBaseURL(fullURL, r.baseURL) {
		blockAfter = true
		return ""
	}
	// recentlyFailed check needs the mutex briefly.
	r.mu.Lock()
	_, recent := r.recentlyFailed[fullURL]
	r.mu.Unlock()
	if recent {
		return ""
	}
	return fullURL
}

// hostMatchesBaseURL returns true iff `rawURL`'s hostname is `baseURL` itself
// or a proper subdomain of it. Used by the discovery step's URL allowlist
// to bound a BFF compromise's blast radius. The leading "." in the suffix
// check rejects "moolabs.com.attacker.com" — without it, a HasSuffix match
// on "moolabs.com" would erroneously accept the attacker host.
func hostMatchesBaseURL(rawURL, baseURL string) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	host := strings.ToLower(parsed.Hostname())
	base := strings.ToLower(strings.TrimLeft(baseURL, "."))
	if host == "" || base == "" {
		return false
	}
	return host == base || strings.HasSuffix(host, "."+base)
}

func (r *IngestURLResolver) regionFallbackURLLocked() string {
	regionCode, ok := RegionIngestMap[r.region]
	if !ok {
		return r.Step4LastResortURL()
	}
	candidate := "https://ingest." + regionCode + "." + r.baseURL + meterIngestPath
	if _, recent := r.recentlyFailed[candidate]; recent {
		return r.Step4LastResortURL()
	}
	return candidate
}

func (r *IngestURLResolver) expireRecentlyFailedLocked() {
	now := r.clock()
	for u, expiry := range r.recentlyFailed {
		if !expiry.After(now) {
			delete(r.recentlyFailed, u)
		}
	}
}
