package logger

import (
	"log"
	"os"
)

func StdoutLogger() *log.Logger {
	return log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
}
