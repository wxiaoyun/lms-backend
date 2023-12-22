package middleware

import (
	"lms-backend/internal/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

var (
	AllowHeaders = []string{
		"Origin",
		"Content-Type",
		"Accept",
		csrf.HeaderName,
	}
)

func SetupCors(app *fiber.App, cfg *config.Config) {
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     cfg.FrontendURL,
		AllowOriginsFunc: allowedOrigins,
		AllowHeaders:     strings.Join(AllowHeaders, ", "),
	}))
}

// Allow anythin from https://*lms-cambodia-dev.netlify.app/
// e.g. https://deploy-preview-21--lms-cambodia-dev.netlify.app/
func allowedOrigins(s string) bool {
	if strings.HasSuffix(s, "lms-cambodia-dev.netlify.app") {
		return true
	}

	if strings.HasSuffix(s, "wxiaoyun.com") {
		return true
	}

	return false
}
