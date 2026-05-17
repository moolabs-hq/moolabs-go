// G5 in-memory ingest buffer — Go port of _dx_buffer.py.
//
// Same shape as Python: bounded queue + background drain worker. Go-specific
// differences:
//   - Goroutine + ticker for periodic drain (idiomatic Go, not threading).
//   - sync.Mutex guards the queue (Python had its own Lock; TS didn't need
//     one since JS is single-threaded — Go does, like Python).
//   - Close() returns an error and blocks up to ShutdownFlushTimeout
//     attempting the final drain.

package moolabs

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Overflow policies. Strings (not enums) so the customer-facing config is
// friendly: IngestBufferConfig{Overflow: moolabs.OverflowDropOldest}.
const (
	OverflowDropOldest = "drop_oldest"
	OverflowRaise      = "raise"
)

// ErrIngestBufferFull is returned by Enqueue when Overflow=OverflowRaise
// and the queue would exceed MaxSize.
var ErrIngestBufferFull = errors.New("moolabs: ingest buffer full")

// IngestBufferConfig tunes the buffer's behavior. Defaults match Python/TS
// (contracts §3.5a / O6).
type IngestBufferConfig struct {
	// MaxSize — events the queue holds before overflow policy fires.
	MaxSize int

	// FlushInterval — background drain tick period.
	FlushInterval time.Duration

	// ShutdownFlushTimeout — Close() blocks at most this long flushing.
	ShutdownFlushTimeout time.Duration

	// Overflow — OverflowDropOldest (default) or OverflowRaise.
	Overflow string

	// Logger receives per-event diagnostic warnings from the buffer
	// (overflow drops, shutdown event abandonment). Defaults to
	// NoopLogger if nil — library never writes to stderr unless the
	// customer explicitly opts in. In normal Moolabs construction the
	// customer's Config.Logger is passed through here automatically.
	Logger Logger
}

// DefaultIngestBufferConfig — same values as Python/TS.
//
// Tunables optimized for per-customer HTTP request rate (3 req/min at
// 20s flush) over per-event latency. Throughput ceiling =
// MaxSize / FlushInterval = 50 evt/s; customers sustaining >25 evt/s
// should raise MaxSize via MoolabsOptions.BufferMax to keep a 2×
// safety margin, otherwise Stats().Dropped will increment continuously
// under normal operation.
//
// Peak memory: ~500 KB worst case (1000 events × ~500 bytes typical).
// Per-event latency to backend: up to FlushInterval (20s).
// Process-exit hang: up to ShutdownFlushTimeout (60s) when the
// backend is unreachable at close() time.
var DefaultIngestBufferConfig = IngestBufferConfig{
	MaxSize:              1000,
	FlushInterval:        20 * time.Second,
	ShutdownFlushTimeout: 60 * time.Second,
	Overflow:             OverflowDropOldest,
}

// DrainCallback is invoked by the background worker. It receives the
// current batch and returns the count of events successfully delivered
// from the FRONT of the batch (so the unsent tail stays queued). Returning
// an error leaves all events queued for next tick.
type DrainCallback func(ctx context.Context, events []any) (int, error)

// IngestBuffer is a bounded in-memory queue with a background drain worker.
type IngestBuffer struct {
	drain  DrainCallback
	config IngestBufferConfig

	mu                      sync.Mutex
	queue                   []any
	enqueued                int
	dropped                 int
	delivered               int
	drainFailures           int
	terminalDrops           int // events dropped because the server returned a non-retryable status
	abandonedOnShutdown     int // events lost when Close() timed out before draining

	wakeup chan struct{} // signals worker to drain now
	done   chan struct{} // signaled when worker exits
	cancel context.CancelFunc
}

// NewIngestBuffer constructs a buffer with the given drain callback.
// Validates the config; returns an error on bad values.
func NewIngestBuffer(drain DrainCallback, cfg IngestBufferConfig) (*IngestBuffer, error) {
	if cfg.MaxSize <= 0 {
		return nil, fmt.Errorf("moolabs: IngestBufferConfig.MaxSize must be positive, got %d", cfg.MaxSize)
	}
	if cfg.FlushInterval <= 0 {
		return nil, fmt.Errorf("moolabs: IngestBufferConfig.FlushInterval must be positive, got %v", cfg.FlushInterval)
	}
	if cfg.ShutdownFlushTimeout < 0 {
		return nil, fmt.Errorf("moolabs: IngestBufferConfig.ShutdownFlushTimeout must be non-negative, got %v", cfg.ShutdownFlushTimeout)
	}
	if cfg.Overflow != OverflowDropOldest && cfg.Overflow != OverflowRaise {
		return nil, fmt.Errorf("moolabs: IngestBufferConfig.Overflow must be %q or %q, got %q",
			OverflowDropOldest, OverflowRaise, cfg.Overflow)
	}
	if drain == nil {
		return nil, errors.New("moolabs: IngestBuffer drain callback cannot be nil")
	}
	if cfg.Logger == nil {
		cfg.Logger = NoopLogger{}   // default = no output
	}
	return &IngestBuffer{
		drain:  drain,
		config: cfg,
		queue:  make([]any, 0, cfg.MaxSize),
		wakeup: make(chan struct{}, 1),
		done:   make(chan struct{}),
	}, nil
}

// Start launches the background drain worker. Idempotent.
func (b *IngestBuffer) Start() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.cancel != nil {
		return // already running
	}
	ctx, cancel := context.WithCancel(context.Background())
	b.cancel = cancel
	go b.workerLoop(ctx)
}

// Enqueue adds events to the queue, honoring the overflow policy.
// Returns ErrIngestBufferFull if Overflow=OverflowRaise and the queue
// would exceed MaxSize.
//
// Note: the customer-supplied Logger.Warn is called AFTER releasing the
// mutex. A logger that takes its own lock (or re-enters the buffer via
// Stats/Qsize) would otherwise deadlock the worker (round 3 review).
func (b *IngestBuffer) Enqueue(events []any) error {
	if len(events) == 0 {
		return nil
	}
	b.mu.Lock()

	newSize := len(b.queue) + len(events)
	overflowed := false
	maxSize := b.config.MaxSize
	if overflow := newSize - maxSize; overflow > 0 {
		if b.config.Overflow == OverflowRaise {
			b.mu.Unlock()
			return fmt.Errorf("%w: at %d events (max=%d); %d rejected",
				ErrIngestBufferFull, len(b.queue), maxSize, len(events))
		}
		// drop_oldest: pop from front to make room.
		popFromExisting := overflow
		if popFromExisting > len(b.queue) {
			popFromExisting = len(b.queue)
		}
		b.queue = b.queue[popFromExisting:]
		b.dropped += popFromExisting
		// If new batch alone exceeds MaxSize, only the tail of size MaxSize fits.
		if len(events) > maxSize {
			excess := len(events) - maxSize
			b.dropped += excess
			events = events[excess:]
		}
		overflowed = true
	}
	b.queue = append(b.queue, events...)
	b.enqueued += len(events)
	// Non-blocking wakeup of the worker (channel has buffer 1; second send is a no-op).
	select {
	case b.wakeup <- struct{}{}:
	default:
	}
	b.mu.Unlock()
	// Logger called OUTSIDE the mutex — see function-level comment.
	if overflowed {
		b.config.Logger.Warn("moolabs.ingestBuffer.overflowDropOldest",
			"maxSize", maxSize)
	}
	return nil
}

// Close stops the worker and attempts one final synchronous drain bounded
// by ShutdownFlushTimeout. Idempotent.
func (b *IngestBuffer) Close() error {
	b.mu.Lock()
	if b.cancel == nil {
		b.mu.Unlock()
		return nil // never started or already closed
	}
	cancel := b.cancel
	b.cancel = nil
	b.mu.Unlock()

	cancel() // signal worker to exit
	// Wait for worker to exit, bounded by timeout
	select {
	case <-b.done:
	case <-time.After(b.config.ShutdownFlushTimeout):
	}
	// Final synchronous drain attempt (best effort)
	ctx, drainCancel := context.WithTimeout(context.Background(), b.config.ShutdownFlushTimeout)
	defer drainCancel()
	_ = b.drainOnce(ctx)
	if leftover := b.Qsize(); leftover > 0 {
		// Bump counter + emit per-event log. Customer with no Logger gets
		// no output (NoopLogger silently discards); customer with a real
		// Logger gets the per-shutdown detail. Counter is always available
		// via Stats().AbandonedOnShutdown for periodic monitoring.
		b.mu.Lock()
		b.abandonedOnShutdown += leftover
		b.mu.Unlock()
		b.config.Logger.Warn("moolabs.ingestBuffer.eventsAbandonedOnShutdown",
			"count", leftover)
	}
	return nil
}

// Qsize returns the current queue length.
func (b *IngestBuffer) Qsize() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.queue)
}

// IngestBufferStats is a snapshot of counters for observability + tests.
//
// All counters are monotonic over the buffer's lifetime. Customers can poll
// Stats() and diff against a previous snapshot to detect silent-failure
// events (TerminalDrops bumping = bad API key; AbandonedOnShutdown bumping
// = Close() timed out). The library does not write to stderr by default —
// pass a Logger via Config.Logger to surface the same events as structured
// log lines as they happen; otherwise polling Stats() is the only signal.
type IngestBufferStats struct {
	Enqueued            int
	Dropped             int // dropped because the buffer was full and overflow policy fired
	Delivered           int
	DrainFailures       int
	TerminalDrops       int // dropped because the server returned a non-retryable status
	AbandonedOnShutdown int // lost because Close() timed out before final drain
}

// Stats returns counter snapshots.
func (b *IngestBuffer) Stats() IngestBufferStats {
	b.mu.Lock()
	defer b.mu.Unlock()
	return IngestBufferStats{
		Enqueued:            b.enqueued,
		Dropped:             b.dropped,
		Delivered:           b.delivered,
		DrainFailures:       b.drainFailures,
		TerminalDrops:       b.terminalDrops,
		AbandonedOnShutdown: b.abandonedOnShutdown,
	}
}

// recordTerminalDrop bumps the terminal-drop counter. Called by Moolabs's
// buffer drain when a non-retryable HTTP status (auth/validation/removed-
// route) is returned for a batch — the events are removed from the queue
// (retrying with the same key/body fails identically), and this counter
// is the SOLE customer-visible signal that data was lost.
func (b *IngestBuffer) recordTerminalDrop(count int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.terminalDrops += count
}

// ── internals ──

func (b *IngestBuffer) workerLoop(ctx context.Context) {
	defer close(b.done)
	ticker := time.NewTicker(b.config.FlushInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		case <-b.wakeup:
		}
		_ = b.drainOnce(ctx)
	}
}

// drainOnce pops the queue into an in-flight batch under the lock BEFORE
// the network call, then re-enqueues any undelivered tail at the FRONT
// after the call returns.
//
// Pre-2026-05-16 design: snapshot-then-slice-after. That had a race where
// concurrent Enqueue with drop_oldest could shift the queue front during
// the drain; the post-drain `b.queue = b.queue[delivered:]` then sliced
// off newly-enqueued events instead of the snapshot. Silent data loss.
//
// New design: pop-everything-under-lock before drain. In-flight events
// are invisible to overflow logic, so concurrent Enqueue with drop_oldest
// can't touch them. Failed delivery → ALL events go back to the front;
// partial delivery → the undelivered tail goes back to the front. Order
// preserved.
//
// Side effect: re-enqueue at front may briefly exceed MaxSize if a
// concurrent Enqueue filled the queue during drain. Next overflow check
// fires on the next Enqueue and the now-oldest (post-drain tail) is the
// first dropped — same end-state as if the failed delivery hadn't happened.
func (b *IngestBuffer) drainOnce(ctx context.Context) error {
	b.mu.Lock()
	if len(b.queue) == 0 {
		b.mu.Unlock()
		return nil
	}
	// Pop ALL events into the in-flight batch under the lock.
	batch := b.queue
	b.queue = nil
	b.mu.Unlock()

	delivered, err := b.drain(ctx, batch)

	// Clamp delivered to batch bounds.
	if delivered < 0 {
		delivered = 0
	}
	if delivered > len(batch) {
		delivered = len(batch)
	}
	undelivered := batch[delivered:]

	b.mu.Lock()
	defer b.mu.Unlock()
	if err != nil {
		b.drainFailures++
	}
	if delivered > 0 {
		b.delivered += delivered
	}
	// Re-enqueue undelivered tail at the FRONT. New events that arrived
	// during the drain go AFTER the re-fronted tail.
	if len(undelivered) > 0 {
		// Concatenate: undelivered ++ queue (re-fronted)
		newQueue := make([]any, 0, len(undelivered)+len(b.queue))
		newQueue = append(newQueue, undelivered...)
		newQueue = append(newQueue, b.queue...)
		b.queue = newQueue
	}
	return err
}

