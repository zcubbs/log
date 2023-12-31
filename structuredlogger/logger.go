package structuredlogger

// StructuredLogger is a generic interface for logging with different loggers.
type StructuredLogger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
	SetFormat(format string)
	SetLevel(level string)
}
