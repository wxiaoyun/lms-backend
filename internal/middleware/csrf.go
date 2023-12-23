package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

const (
	MaxAge = time.Hour * 1 // 1 hour
)

var CSRFMiddleware = csrf.New(
	csrf.Config{
		KeyLookup:      "header:" + csrf.HeaderName,
		CookieName:     "__Secure-csrf_",
		CookieSameSite: "None",
		CookieSecure:   true,
		CookieHTTPOnly: false,
		Expiration:     MaxAge,
		// Session:           session.Store,
	},
)

func SetupCSRF(app *fiber.App) {
	app.Use(csrf.New(
	// csrf.Config{
	// 	KeyLookup:         "header:" + csrf.HeaderName,
	// 	CookieName:        "__Secure-csrf_",
	// 	CookieSameSite:    "None",
	// 	CookieSecure:      true,
	// 	CookieSessionOnly: true,
	// 	CookieHTTPOnly:    false,
	// 	CookieDomain:      domain(cfg),
	// 	Expiration:        MaxAge,
	// 	Extractor:         csrf.CsrfFromHeader(csrf.HeaderName),
	// 	// Session:           session.Store,
	// },
	))
}

// func domain(cfg *config.Config) string {
// 	if strings.HasSuffix(cfg.FrontendURL, "wxiaoyun.com") {
// 		return ".wxiaoyun.com"
// 	}

// 	return ""
// }
