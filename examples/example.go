package main

import (
	"github.com/zcubbs/log"
	"github.com/zcubbs/log/structuredlogger"
)

func main() {
	// Creating a new Logrus logger in JSON format
	logrusLogger := log.NewLogger(structuredlogger.LogrusLoggerType, "LogrusLogger", structuredlogger.JSONFormat)
	logrusLogger.Info("This is a Logrus logger in JSON format")

	// Switching the Logrus logger to text format
	logrusLogger.SetFormat(structuredlogger.TextFormat)
	logrusLogger.Info("This is a Logrus logger in text format")

	// Creating a new Zap logger in JSON format
	zapLogger := log.NewLogger(structuredlogger.ZapLoggerType, "ZapLogger", structuredlogger.JSONFormat)
	zapLogger.Info("This is a Zap logger in JSON format")

	// Switching the Zap logger to text format
	zapLogger.SetFormat(structuredlogger.TextFormat)
	zapLogger.Info("This is a Zap logger in text format")

	// Creating a new Standard logger in JSON format
	stdLogger := log.NewLogger(structuredlogger.StdLoggerType, "StdLogger", structuredlogger.JSONFormat)
	stdLogger.Info("This is a Standard logger in JSON format")

	// Switching the Standard logger to text format
	stdLogger.SetFormat(structuredlogger.TextFormat)
	stdLogger.Info("This is a Standard logger in text format")
}
