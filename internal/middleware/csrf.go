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
