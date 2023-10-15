package logwrapper

import (
	"github.com/zcubbs/logwrapper/logger"
	"github.com/zcubbs/logwrapper/loggers"
)

// NewLogger is a factory function that creates a new logger based on the given type.
func NewLogger(loggerType string, name string, format string) *logger.Logger {
	l := logger.Logger(nil)

	switch loggerType {
	case logger.LogrusLoggerType:
		l = loggers.NewLogrusLogger(name)
	case logger.ZapLoggerType:
		l = loggers.NewZapLogger(name)
	case logger.CharmLoggerType:
		l = loggers.NewCharmLogger(name)
	case logger.StdLoggerType:
		l = loggers.NewStdLogger(name)
	default:
		l = loggers.NewStdLogger(name)
	}

	l.SetFormat(format)
	return &l
}
