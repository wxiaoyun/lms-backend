package helper

import (
	"math/rand"
	"strings"
	"time"

	"github.com/go-loremipsum/loremipsum"
)

var loremIpsumGenerator = loremipsum.New()

func RandInt(from, to int) int {
	// //nolint:gosec
	return rand.Intn(to-from) + from
}

func RandomDate(start, end time.Time) time.Time {
	min := start.Unix()
	max := end.Unix()
	delta := max - min

	// //nolint:gosec
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func RandWords(length int) string {
	builder := strings.Builder{}

	for i := 0; i < length; i++ {
		if _, err := builder.WriteString(loremIpsumGenerator.Word() + " "); err != nil {
			return "error has occurred generating random words"
		}
	}

	return builder.String()
}
