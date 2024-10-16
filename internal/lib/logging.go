package lib

import (
	"log"
	"os"
)

type LogManager struct {
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
}

var Logger *LogManager

func NewLogger() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Logger = &LogManager{
		Info:    log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Warning: log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error:   log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
