package loggers

import (
	"github.com/zcubbs/logwrapper/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger(name string) logger.Logger {
	zapLogger := &ZapLogger{}
	zapLogger.initLogger(name, logger.TextFormat) // default to text format
	return zapLogger
}

func (l *ZapLogger) initLogger(name string, format string) {
	var zapLogger *zap.Logger
	var err error

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string{"stdout"}

	switch format {
	case logger.JSONFormat:
		cfg.Encoding = "json"
	case logger.TextFormat:
		cfg.Encoding = "console"
	}

	zapLogger, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	l.logger = zapLogger.Sugar().Named(name)
}

func (l *ZapLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *ZapLogger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *ZapLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (l *ZapLogger) SetFormat(format string) {
	l.initLogger(l.logger.Desugar().Name(), format)
}

func (l *ZapLogger) SetLevel(level string) {
	var zapLevel zapcore.Level
	switch level {
	case logger.DebugLevel:
		zapLevel = zapcore.DebugLevel
	case logger.InfoLevel:
		zapLevel = zapcore.InfoLevel
	case logger.WarnLevel:
		zapLevel = zapcore.WarnLevel
	case logger.ErrorLevel:
		zapLevel = zapcore.ErrorLevel
	case logger.FatalLevel:
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.InfoLevel
	}
	l.logger.Desugar().Core().Enabled(zapLevel)
}
