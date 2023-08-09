package helper

import (
	"fmt"
	"math/rand"
	"time"
)

// generateISBN13 generates a random ISBN-13 number.
func GenerateISBN13() string {
	rand.Seed(time.Now().UnixNano())

	// Prefix
	prefix := 978

	// Group, publisher, and item numbers
	//nolint:gosec
	groupPublisherItem := rand.Intn(1000000000)

	// Compute the checksum
	isbn := fmt.Sprintf("%d%09d", prefix, groupPublisherItem)
	checksum := 0
	for i, digit := range isbn {
		d := int(digit - '0')
		if i%2 == 0 {
			checksum += d
		} else {
			checksum += 3 * d
		}
	}
	checksum = (10 - (checksum % 10)) % 10

	// Combine all parts to create the ISBN-13
	return fmt.Sprintf("%d%09d%d", prefix, groupPublisherItem, checksum)
}
