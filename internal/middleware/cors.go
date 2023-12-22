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
		AllowHeaders:     strings.Join(AllowHeaders, ", "),
	}))
}
