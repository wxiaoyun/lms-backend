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
		AllowOriginsFunc: allowedOrigins(cfg),
		AllowHeaders:     strings.Join(AllowHeaders, ", "),
	}))
}

func allowedOrigins(cfg *config.Config) func(string) bool {
	return func(header string) bool {
		if strings.HasSuffix(header, cfg.FrontendURL) {
			return true
		}

		if cfg.Mode == "development" {
			return true
		}

		return false
	}
}
