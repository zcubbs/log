package loggers

import (
	"encoding/json"
	"fmt"
	"github.com/zcubbs/logwrapper/logger"
	"log"
	"os"
)

type StdLogger struct {
	format string
	name   string
}

func NewStdLogger(name string) logger.Logger {
	l := &StdLogger{name: name}
	l.SetFormat(logger.TextFormat) // setting default to TextFormat
	return l
}

func (s *StdLogger) SetFormat(format string) {
	s.format = format
}

func (s *StdLogger) logMessage(level string, msg string, keysAndValues ...interface{}) {
	if s.format == logger.JSONFormat {
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
		log.Println(string(jsonMessage))
	} else {
		log.Printf("%s [%s]: %s %v", level, s.name, msg, keysAndValues)
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
