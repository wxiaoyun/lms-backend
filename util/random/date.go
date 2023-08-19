package random

import (
	"math/rand"
	"time"
)

func RandomDate(start, end time.Time) time.Time {
	min := start.Unix()
	max := end.Unix()
	delta := max - min

	// //nolint:gosec
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
