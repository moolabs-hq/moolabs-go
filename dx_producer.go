// Package moolabs — producer-channel pattern.
//
// Extracted from dx_client.go (round-5 refactor — round-4 review I-NEW-2)
// so the producer logic is leaf-testable: no dependency on the openapi-
// generator-generated types (EventsAPIService, APIClient, Configuration).
// UsageNamespace embeds *ProducerChannel and delegates IngestEvents'
// non-blocking path to it.
//
// What this gives us:
//   - Customer's IngestEvents call does ONLY a non-blocking channel send
//     (~50 ns). It never touches the buffer's mutex.
//   - A dedicated producer goroutine drains the channel and invokes a
//     caller-supplied delegate (in production: buffer.Enqueue).
//   - Panic recovery: a buggy customer Logger that panics during
//     delegate execution does NOT kill the producer goroutine; it
//     bumps producerPanics and respawns.
//   - Channel capacity is configurable (in production: max(1024, BufferMax/8)).
//   - Lifecycle: lazy start via sync.Once; explicit Stop() drains the
//     channel before exit.
//
// The delegate function is the seam — tests pass a recording or
// counting delegate instead of a real buffer.Enqueue.

package moolabs

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// ProducerChannel decouples the customer's non-blocking submission from
// the buffer's mutex-bound enqueue.
//
//	customer goroutine ──Submit→ [channel] ──producerLoop→ delegate
//
// Submit is non-blocking — on channel full it returns false and increments
// ingestQueueDropped (observable via IngestQueueDropped()).
type ProducerChannel struct {
	// eventCh is the producer-channel. Customers write batches; the
	// producer goroutine reads them.
	eventCh chan []any

	// delegate is called for each batch the producer reads from eventCh.
	// In production: u.buffer.Enqueue. In tests: a recording function.
	// MUST be safe to call concurrently with Submit + Stop.
	delegate func([]any)

	// logger receives panic-recovery notifications; nil-safe.
	logger Logger

	// chanCap is the channel buffer size. Configured at construction.
	chanCap int

	// Lifecycle
	producerOnce sync.Once
	producerStop chan struct{}
	producerDone chan struct{}

	// Counters
	ingestQueueDropped atomic.Int64
	producerPanics     atomic.Int64
}

// NewProducerChannel constructs a producer-channel ready for Submit.
// The goroutine is started lazily on first Submit (or explicit Start()).
//
// chanCap MUST be positive. logger may be nil (will treat as noop).
// delegate MUST NOT be nil.
func NewProducerChannel(chanCap int, delegate func([]any), logger Logger) *ProducerChannel {
	if chanCap <= 0 {
		panic("moolabs: ProducerChannel chanCap must be positive")
	}
	if delegate == nil {
		panic("moolabs: ProducerChannel delegate cannot be nil")
	}
	if logger == nil {
		logger = NoopLogger{}
	}
	return &ProducerChannel{
		delegate: delegate,
		logger:   logger,
		chanCap:  chanCap,
	}
}

// Submit attempts to enqueue events on the channel. Returns true if the
// channel accepted them; false if the channel was full (events dropped,
// counter bumped).
//
// Non-blocking: ~50 ns on the customer's goroutine. Never holds any
// mutex; never does HTTP; never calls the delegate.
//
// On first call, the producer goroutine is started lazily via sync.Once.
func (p *ProducerChannel) Submit(events []any) bool {
	p.ensureStarted()
	select {
	case p.eventCh <- events:
		return true
	default:
		p.ingestQueueDropped.Add(int64(len(events)))
		return false
	}
}

// Start eagerly starts the producer goroutine. Idempotent — Submit
// calls Start lazily on first invocation. Exposed so callers that want
// the goroutine running BEFORE any Submit can ensure it (e.g., warm-up
// at SDK initialization).
func (p *ProducerChannel) Start() {
	p.ensureStarted()
}

func (p *ProducerChannel) ensureStarted() {
	p.producerOnce.Do(func() {
		p.eventCh = make(chan []any, p.chanCap)
		p.producerStop = make(chan struct{})
		p.producerDone = make(chan struct{})
		go p.loop()
	})
}

// loop drains eventCh and calls delegate per batch. Wrapped in
// defer/recover so a panicking delegate (e.g., customer Logger nil-deref
// surfaced from the buffer's enqueue) does NOT kill the producer.
// On panic: bump producerPanics, log via injected Logger, respawn.
func (p *ProducerChannel) loop() {
	defer func() {
		if r := recover(); r != nil {
			p.producerPanics.Add(1)
			// Log via injected Logger. Defensively swallow logger
			// panics — don't let the recovery itself crash.
			func() {
				defer func() { _ = recover() }()
				p.logger.Warn("moolabs.producer.recovered_from_panic",
					"panic", fmt.Sprint(r))
			}()
			// Restart loop in a fresh goroutine. We did NOT close
			// producerDone (the unrecovered defer was suppressed),
			// so the new loop's defer will close it on exit.
			go p.loop()
			return
		}
		close(p.producerDone)
	}()
	for {
		select {
		case events := <-p.eventCh:
			p.delegate(events)
		case <-p.producerStop:
			// Drain any remaining buffered events before exiting.
			for {
				select {
				case events := <-p.eventCh:
					p.delegate(events)
				default:
					return
				}
			}
		}
	}
}

// Stop signals the producer to exit and waits for it to finish draining
// the channel. Idempotent — safe to call before Start (no-op).
//
// After Stop returns, no further delegate calls will fire (in-flight
// delegate calls complete normally before the loop checks for more work).
func (p *ProducerChannel) Stop() {
	if p.producerStop == nil {
		return // never started
	}
	// Idempotent close via a sync.Once-like pattern (close panics if
	// double-closed; recover swallows).
	func() {
		defer func() { _ = recover() }()
		close(p.producerStop)
	}()
	<-p.producerDone
}

// SubmitAndCount is the typed contract IngestEvents uses to translate
// Submit's boolean into a "delivered count" value. Returns len(events)
// when the producer accepted the batch, 0 when the channel was full
// (events dropped, counter bumped).
//
// Codex round-6 HIGH-2 introduced this helper so the count semantics
// are formalized in one place — pre-fix, IngestEvents discarded the
// Submit boolean and always returned len(events), lying to the caller
// about acceptance on a full channel. Tests pin the contract here
// (TestProducerChannel_SubmitAndCount_*), making the regression
// impossible to silently re-introduce: the obvious call shape becomes
// `producer.SubmitAndCount(events)`, not `producer.Submit(events)`.
func (p *ProducerChannel) SubmitAndCount(events []any) int {
	if p.Submit(events) {
		return len(events)
	}
	return 0
}

// IngestQueueDropped returns the count of events dropped because the
// channel was full at the moment of Submit (caller could not enqueue).
// Distinct from buffer-layer drop_oldest counters.
func (p *ProducerChannel) IngestQueueDropped() int64 {
	return p.ingestQueueDropped.Load()
}

// ProducerPanics returns the count of times the producer goroutine
// recovered from a panic in the delegate. A non-zero value indicates
// either a buggy customer Logger or an unexpected delegate failure —
// the SDK is still functional (loop respawned) but operators should
// audit the delegate path for bugs.
func (p *ProducerChannel) ProducerPanics() int64 {
	return p.producerPanics.Load()
}
