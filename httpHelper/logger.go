package httpHelper

import (
	"log"
	"os"
)

// InfoLog and ErrorLog represent logging instances for informational and error messages.
var (
	InfoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
)

// SetInfoLogger sets a custom logger for informational messages.
func SetInfoLogger(log *log.Logger) {
	InfoLog = log
}

// SetErrorLogger sets a custom logger for error messages.
func SetErrorLogger(log *log.Logger) {
	ErrorLog = log
}
