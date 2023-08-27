package middleware

import (
	sessionmiddleware "lms-backend/internal/middleware/session"
	"lms-backend/internal/session"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// attach middleware
func SetupAppMiddleware(app *fiber.App) {
	session.SetupStore()

	app.Use(sessionmiddleware.SessionMiddleware)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     os.Getenv("FRONTEND_URL"),
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))
}
