package middleware

import (
	"lms-backend/internal/session"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

const (
	MaxAge         = time.Hour * 1 // 1 hour
	CSRFContextKey = "csrf"
)

func SetupCSRF(app *fiber.App) {
	app.Use(csrf.New(csrf.Config{
		KeyLookup:         "header:" + csrf.HeaderName,
		CookieName:        "__Host-csrf_",
		CookieSameSite:    "Strict",
		CookieSecure:      true,
		CookieHTTPOnly:    true,
		CookieSessionOnly: true,
		ContextKey:        CSRFContextKey,
		Expiration:        MaxAge,
		KeyGenerator:      utils.UUIDv4,
		Extractor:         csrf.CsrfFromHeader(csrf.HeaderName),
		Session:           session.Store,
		SessionKey:        "fiber.csrf.token",
		HandlerContextKey: "fiber.csrf.handler",
	}))
}
