package logs

import (
	//
	"os"
	"time"

	//
	"github.com/kataras/golog"
)

var logger *golog.Logger

// Initial : create Instance Logger
func Initial() *golog.Logger {
	logger = golog.New()
	logger.SetOutput(os.Stdout)
	logger.SetTimeFormat(time.RFC3339)
	logger.SetLevel(os.Getenv("LOG_LEVEL"))
	logger.Debug("Logger: create Instance")
	return logger
}

// Logger : get Logger
func Logger() *golog.Logger {
	if logger == nil {
		logger = Initial()
	}
	return logger
}
