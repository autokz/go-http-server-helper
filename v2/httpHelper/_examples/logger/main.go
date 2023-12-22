package main

import (
	"github.com/autokz/go-http-server-helper/v2/httpHelper"
	"log"
	"os"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	httpHelper2.SetInfoLogger(infoLog)
	httpHelper2.SetErrorLogger(errorLog)
}
