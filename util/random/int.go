package random

import (
	"math/rand"
)

func RandInt(from, to int) int {
	// //nolint:gosec
	return rand.Intn(to-from) + from
}
