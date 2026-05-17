// Package logger — minimal opt-in logging surface for the Moolabs SDK.
//
// The customer-facing contract: by default the SDK never writes to stderr.
// If the customer wants per-event diagnostics, they pass a Logger via
// Config.Logger; if not, NoopLogger is used and the library produces no
// output.
//
// Counters on IngestBufferStats remain the primary monitoring signal
// (aggregable, polling-friendly). Logger calls give per-event context
// (status code, raw error, batch size) that the counter alone can't
// provide — useful when a customer is debugging a specific failure.
//
// This file is intentionally tiny and dependency-free so it can be
// compiled with the leaf-only test invocation (which excludes
// dx_client.go because dx_client.go pulls in openapi-generator-cli's
// generated types).

package moolabs

// Logger is the minimal interface the SDK uses for per-event diagnostics.
// Implementations route the message + structured fields to the customer's
// preferred logging stack (zap, zerolog, slog, etc).
//
// Default behavior when Config.Logger is nil is NoopLogger{} — no output.
type Logger interface {
	// Warn emits a per-event warning. `msg` is a stable identifier
	// (e.g. "moolabs.ingest_buffer.terminal_drop"); `kvs` is alternating
	// (string-key, value, string-key, value, ...). Implementations may
	// panic on odd-length kvs — callers always pass even-length.
	Warn(msg string, kvs ...any)
}

// NoopLogger is the default Logger — discards every call.
// Used when Config.Logger is nil. Library never writes to stderr.
type NoopLogger struct{}

// Warn discards its arguments.
func (NoopLogger) Warn(_ string, _ ...any) {}
