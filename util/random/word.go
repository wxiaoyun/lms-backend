package random

import (
	"strings"

	"github.com/go-loremipsum/loremipsum"
)

var loremIpsumGenerator = loremipsum.New()

func RandWords(length int) string {
	builder := strings.Builder{}

	for i := 0; i < length; i++ {
		if _, err := builder.WriteString(loremIpsumGenerator.Word() + " "); err != nil {
			return "error has occurred generating random words"
		}
	}

	return builder.String()
}
