package lib

import (
	"io"
	"log"
	"os"
)

type LogManager struct {
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
}

var (
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
)

func NewLogger(log_level string) {
	logger := &LogManager{
		ErrorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	switch log_level {
	case "debug":
		logger.InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	case "warnings":
		logger.InfoLogger = log.New(io.Discard, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	case "errors":
		logger.InfoLogger = log.New(io.Discard, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.WarningLogger = log.New(io.Discard, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	LogInfo = logger.InfoLogger
	LogWarning = logger.WarningLogger
	LogError = logger.ErrorLogger
}
