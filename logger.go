package log

import (
	"github.com/zcubbs/log/loggers"
	"github.com/zcubbs/log/structuredlogger"
)

func init() {
	loggerName = "StdLogger"
	loggerFormat = structuredlogger.TextFormat
	loggrType = structuredlogger.StdLoggerType
	loggerLevel = structuredlogger.InfoLevel

	l := NewLogger(loggrType, loggerName, loggerFormat)
	l.SetLevel(loggerLevel)
	logw = &l
}

var (
	logw         *Logger
	loggrType    string
	loggerName   string
	loggerFormat string
	loggerLevel  string
)

type Logger interface {
	structuredlogger.StructuredLogger
}

// NewLogger is a factory function that creates a new logger based on the given type.
func NewLogger(loggerType string, name string, format string) Logger {
	l := structuredlogger.StructuredLogger(nil)

	switch loggerType {
	case structuredlogger.LogrusLoggerType:
		l = loggers.NewLogrusLogger(name)
	case structuredlogger.ZapLoggerType:
		l = loggers.NewZapLogger(name)
	case structuredlogger.CharmLoggerType:
		l = loggers.NewCharmLogger(name)
	case structuredlogger.StdLoggerType:
		l = loggers.NewStdLogger(name)
	default:
		l = loggers.NewStdLogger(name)
	}

	l.SetFormat(format)
	l.SetLevel(structuredlogger.InfoLevel)

	return l
}

func GetLogger() Logger {
	return *logw
}

func SetLogger(logger Logger) {
	logw = &logger
}

func Info(msg string, keysAndValues ...interface{}) {
	GetLogger().Info(msg, keysAndValues...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	GetLogger().Debug(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	GetLogger().Warn(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	GetLogger().Error(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	GetLogger().Fatal(msg, keysAndValues...)
}

func SetLevel(level string) {
	loggerLevel = level
	GetLogger().SetLevel(loggerLevel)
}

func SetFormat(format string) {
	loggerFormat = format
	GetLogger().SetFormat(loggerFormat)
}

func SetLoggerType(loggerType string) {
	loggrType = loggerType

	l := NewLogger(loggrType, loggerName, loggerFormat)
	l.SetLevel(loggerLevel)

	logw = &l
}
