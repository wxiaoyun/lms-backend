// nolint
package util

import (
	"fmt"
)

// Debug is a wrapper around fmt.Print() to make it easier to find all debug statements
//
// It surrounds what to be printed with many line breaks to make it easier to find
//
//nolint:predeclared // ignore error
func Debug(a ...any) {
	fmt.Print("\n\n\n")
	fmt.Print(a...)
	fmt.Print("\n\n\n")
}
