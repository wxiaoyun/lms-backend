package util

import (
	"fmt"
)

// Debug is a wrapper around fmt.Print() to make it easier to find all debug statements
//
// It surrounds what to be printed with many line breaks to make it easier to find
//
//nolint:predeclared // ignore error
func Debug(any ...any) {
	//nolint:revive // ignore error
	fmt.Print("\n\n\n")
	//nolint:revive // ignore error
	fmt.Print(any...)
	//nolint:revive // ignore error
	fmt.Print("\n\n\n")
}
