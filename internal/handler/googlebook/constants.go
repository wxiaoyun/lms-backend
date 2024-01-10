package googlebook

import (
	"os"
)

var (
	// GoogleBookAPIKey is the API key for Google Books API
	googleAPIKey = ""
)

const (
	googleAPIBaseURL = "https://www.googleapis.com/books/v1"
)

func GetGoogleAPIKey() string {
	if googleAPIKey == "" {
		googleAPIKey = os.Getenv("GOOGLE_API_KEY")
	}

	return googleAPIKey
}
