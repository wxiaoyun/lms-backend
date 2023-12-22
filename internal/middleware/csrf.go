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
	app.Use(csrf.New(
		csrf.Config{
			KeyLookup:         "header:" + csrf.HeaderName,
			CookieName:        "__Host-csrf_",
			CookieSameSite:    "None",
			CookieSecure:      true,
			CookieSessionOnly: true,
			CookieHTTPOnly:    false,
			Expiration:        MaxAge,
			Extractor:         csrf.CsrfFromHeader(csrf.HeaderName),
			Session:           session.Store,
			SessionKey:        "fiber.csrf.token",
			HandlerContextKey: "fiber.csrf.handler",
		}))
}
