package middleware

import (
	sessionmiddleware "technical-test/internal/middleware/session"
	"technical-test/internal/session"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// attach middleware
func SetupAppMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	session.SetupStore()

	app.Use(sessionmiddleware.SessionMiddleware)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
}
