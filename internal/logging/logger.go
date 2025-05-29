package logging

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
)

func NewLogger(level LogLevel) *log.Logger {
	logger := log.New()
	
	switch strings.ToLower(string(level)) {
	case string(LevelDebug):
		logger.SetLevel(log.DebugLevel)
	case string(LevelInfo):
		logger.SetLevel(log.InfoLevel)
	case string(LevelWarn):
		logger.SetLevel(log.WarnLevel)
	case string(LevelError):
		logger.SetLevel(log.ErrorLevel)
	default:
		logger.SetLevel(log.InfoLevel)
	}

	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	return logger
}
