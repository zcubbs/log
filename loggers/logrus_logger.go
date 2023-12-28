package loggers

import (
	logruslog "github.com/sirupsen/logrus"
	"github.com/zcubbs/log/structuredlogger"
)

type LogrusLogger struct {
	*logruslog.Entry
	baseLogger *logruslog.Logger
}

func NewLogrusLogger(name string) structuredlogger.StructuredLogger {
	l := logruslog.New()
	l.SetFormatter(&logruslog.JSONFormatter{})
	return &LogrusLogger{l.WithField("logger_name", name), l}
}

func (l *LogrusLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.Entry.Debugf(msg, keysAndValues...)
}

func (l *LogrusLogger) Info(msg string, keysAndValues ...interface{}) {
	l.Entry.Infof(msg, keysAndValues...)
}

func (l *LogrusLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.Entry.Warnf(msg, keysAndValues...)
}

func (l *LogrusLogger) Error(msg string, keysAndValues ...interface{}) {
	l.Entry.Errorf(msg, keysAndValues...)
}

func (l *LogrusLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.Entry.Fatalf(msg, keysAndValues...)
}

func (l *LogrusLogger) SetFormat(format string) {
	switch format {
	case structuredlogger.JSONFormat:
		l.baseLogger.SetFormatter(&logruslog.JSONFormatter{})
	default:
		l.baseLogger.SetFormatter(&logruslog.TextFormatter{})
	}
}

func (l *LogrusLogger) SetLevel(level string) {
	switch level {
	case structuredlogger.DebugLevel:
		l.baseLogger.SetLevel(logruslog.DebugLevel)
	case structuredlogger.InfoLevel:
		l.baseLogger.SetLevel(logruslog.InfoLevel)
	case structuredlogger.WarnLevel:
		l.baseLogger.SetLevel(logruslog.WarnLevel)
	case structuredlogger.ErrorLevel:
		l.baseLogger.SetLevel(logruslog.ErrorLevel)
	}
}
