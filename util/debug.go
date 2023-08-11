package util

import (
	logger "lms-backend/internal/log"
)

var lgr = logger.StdoutLogger()

// Debug is a wrapper around fmt.Print() to make it easier to find all debug statements
//
// It surrounds what to be printed with many line breaks to make it easier to find
//
//nolint:predeclared // ignore error
func Debug(a ...any) {
	lgr.Print("\n\n\n")
	lgr.Print(a...)
	lgr.Print("\n\n\n")
}
