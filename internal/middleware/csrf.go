package middleware

import (
	"lms-backend/internal/session"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

const (
	MaxAge = time.Hour * 1 // 1 hour
)

// fiber csrf middleware cannot be used as it enforces referer and host to be the same.
// since we are hosting backend and frontend on different domains, we need to use custom csrf middleware.
func SetupCSRF(app *fiber.App) {
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:" + csrf.HeaderName,
		CookieName:     "__Secure-csrf_",
		CookieSameSite: "Strict",
		CookieSecure:   true,
		CookieHTTPOnly: false,
		Expiration:     MaxAge,
		Session:        session.Store,
	}))
}
