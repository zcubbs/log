package loggers

import (
	"encoding/json"
	"fmt"
	"github.com/zcubbs/log/structuredlogger"
	stdlog "log"
	"os"
)

type StdLogger struct {
	format string
	name   string
}

func NewStdLogger(name string) structuredlogger.StructuredLogger {
	l := &StdLogger{name: name}
	l.SetFormat(structuredlogger.TextFormat) // setting default to TextFormat
	return l
}

func (s *StdLogger) SetFormat(format string) {
	s.format = format
}

func (s *StdLogger) logMessage(level string, msg string, keysAndValues ...interface{}) {
	if s.format == structuredlogger.JSONFormat {
		logMessage := map[string]interface{}{
			"level": level,
			"name":  s.name,
			"msg":   msg,
		}
		for i := 0; i < len(keysAndValues); i += 2 {
			key := fmt.Sprintf("%v", keysAndValues[i])
			logMessage[key] = keysAndValues[i+1]
		}
		jsonMessage, _ := json.Marshal(logMessage)
		stdlog.Println(string(jsonMessage))
	} else {
		stdlog.Printf("%s [%s]: %s %v", level, s.name, msg, keysAndValues)
	}
}

func (s *StdLogger) Debug(msg string, keysAndValues ...interface{}) {
	s.logMessage("DEBUG", msg, keysAndValues...)
}

func (s *StdLogger) Info(msg string, keysAndValues ...interface{}) {
	s.logMessage("INFO", msg, keysAndValues...)
}

func (s *StdLogger) Warn(msg string, keysAndValues ...interface{}) {
	s.logMessage("WARN", msg, keysAndValues...)
}

func (s *StdLogger) Error(msg string, keysAndValues ...interface{}) {
	s.logMessage("ERROR", msg, keysAndValues...)
}

func (s *StdLogger) Fatal(msg string, keysAndValues ...interface{}) {
	s.logMessage("FATAL", msg, keysAndValues...)
	os.Exit(1)
}

func (s *StdLogger) SetLevel(_ string) {
	// noop
}
