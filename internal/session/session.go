package session

import (
	"lms-backend/internal/database"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	Store *session.Store
)

const (
	CookieKey = "token"
	UserIDKey = "UserID"
	MaxAge    = time.Hour * 24 * 7 // 7 days
)

func SetupStore() {
	Store = session.New(session.Config{
		KeyLookup:      "cookie:token",
		CookieSameSite: "Strict",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		Expiration:     MaxAge,
		Storage:        database.GetRedisStore(),
	})
}
