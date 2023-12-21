package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	Store *session.Store
)

const (
	CookieKey = "token"
	MaxAge    = time.Hour * 24 * 7 // 7 days
	UserIDKey = "UserID"
)

func SetupStore() {
	Store = session.New(session.Config{
		Expiration:     MaxAge,
		KeyLookup:      "cookie:token",
		CookieSameSite: "None",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookiePath:     "/",
	})
}
