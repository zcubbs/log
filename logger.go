package log

import (
	"github.com/zcubbs/log/loggers"
	"github.com/zcubbs/log/structuredlogger"
)

func init() {
	loggerName = "default"
	loggerFormat = structuredlogger.JSONFormat
	loggrType = structuredlogger.CharmLoggerType
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
	switch (*logw).(type) {
	case *loggers.ZapLogger:
		(*logw).(*loggers.ZapLogger).Info(msg, keysAndValues...)
	case *loggers.LogrusLogger:
		(*logw).(*loggers.LogrusLogger).Info(msg, keysAndValues...)
	case *loggers.CharmLogger:
		(*logw).(*loggers.CharmLogger).Info(msg, keysAndValues...)
	case *loggers.StdLogger:
		(*logw).(*loggers.StdLogger).Info(msg, keysAndValues...)
	}
}

func Debug(msg string, keysAndValues ...interface{}) {
	switch (*logw).(type) {
	case *loggers.ZapLogger:
		(*logw).(*loggers.ZapLogger).Debug(msg, keysAndValues...)
	case *loggers.LogrusLogger:
		(*logw).(*loggers.LogrusLogger).Debug(msg, keysAndValues...)
	case *loggers.CharmLogger:
		(*logw).(*loggers.CharmLogger).Debug(msg, keysAndValues...)
	case *loggers.StdLogger:
		(*logw).(*loggers.StdLogger).Debug(msg, keysAndValues...)
	}
}

func Warn(msg string, keysAndValues ...interface{}) {
	switch (*logw).(type) {
	case *loggers.ZapLogger:
		(*logw).(*loggers.ZapLogger).Warn(msg, keysAndValues...)
	case *loggers.LogrusLogger:
		(*logw).(*loggers.LogrusLogger).Warn(msg, keysAndValues...)
	case *loggers.CharmLogger:
		(*logw).(*loggers.CharmLogger).Warn(msg, keysAndValues...)
	case *loggers.StdLogger:
		(*logw).(*loggers.StdLogger).Warn(msg, keysAndValues...)
	}
}

func Error(msg string, keysAndValues ...interface{}) {
	switch (*logw).(type) {
	case *loggers.ZapLogger:
		(*logw).(*loggers.ZapLogger).Error(msg, keysAndValues...)
	case *loggers.LogrusLogger:
		(*logw).(*loggers.LogrusLogger).Error(msg, keysAndValues...)
	case *loggers.CharmLogger:
		(*logw).(*loggers.CharmLogger).Error(msg, keysAndValues...)
	case *loggers.StdLogger:
		(*logw).(*loggers.StdLogger).Error(msg, keysAndValues...)
	}
}

func Fatal(msg string, keysAndValues ...interface{}) {
	switch (*logw).(type) {
	case *loggers.ZapLogger:
		(*logw).(*loggers.ZapLogger).Fatal(msg, keysAndValues...)
	case *loggers.LogrusLogger:
		(*logw).(*loggers.LogrusLogger).Fatal(msg, keysAndValues...)
	case *loggers.CharmLogger:
		(*logw).(*loggers.CharmLogger).Fatal(msg, keysAndValues...)
	case *loggers.StdLogger:
		(*logw).(*loggers.StdLogger).Fatal(msg, keysAndValues...)
	}
}

func SetLevel(level string) {
	loggerLevel = level
	switch (*logw).(type) {
	case *loggers.ZapLogger:
		(*logw).(*loggers.ZapLogger).SetLevel(level)
	case *loggers.LogrusLogger:
		(*logw).(*loggers.LogrusLogger).SetLevel(level)
	case *loggers.CharmLogger:
		(*logw).(*loggers.CharmLogger).SetLevel(level)
	case *loggers.StdLogger:
		(*logw).(*loggers.StdLogger).SetLevel(level)
	}
}

func GetLevel() string {
	return loggerLevel
}

func SetFormat(format string) {
	loggerFormat = format
	switch (*logw).(type) {
	case *loggers.ZapLogger:
		(*logw).(*loggers.ZapLogger).SetFormat(format)
	case *loggers.LogrusLogger:
		(*logw).(*loggers.LogrusLogger).SetFormat(format)
	case *loggers.CharmLogger:
		(*logw).(*loggers.CharmLogger).SetFormat(format)
	case *loggers.StdLogger:
		(*logw).(*loggers.StdLogger).SetFormat(format)
	}
}

func SetLoggerType(loggerType string) {
	loggrType = loggerType

	l := NewLogger(loggrType, loggerName, loggerFormat)
	l.SetLevel(loggerLevel)

	logw = &l
}

func GetLoggerType() string {
	return loggrType
}
