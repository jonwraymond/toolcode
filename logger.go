package toolcode

// Logger is an optional interface for observability during code execution.
// Implementations can log tool calls, timing information, and other events.
type Logger interface {
	// Logf logs a formatted message.
	Logf(format string, args ...any)
}
