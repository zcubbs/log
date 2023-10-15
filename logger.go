package logwrapper

import "github.com/zcubbs/logwrapper/loggers"

// Logger is a generic interface for logging with different loggers.
type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
	SetFormat(format string)
}

// NewLogger is a factory function that creates a new logger based on the given type.
func NewLogger(loggerType string, name string, format string) Logger {
	logger := Logger(nil)

	switch loggerType {
	case LogrusLoggerType:
		logger = loggers.NewLogrusLogger(name)
	case ZapLoggerType:
		logger = loggers.NewZapLogger(name)
	case CharmLoggerType:
		logger = loggers.NewCharmLogger(name)
	case StdLoggerType:
		logger = loggers.NewStdLogger(name)
	default:
		logger = loggers.NewStdLogger(name)
	}

	logger.SetFormat(format)
	return logger
}
