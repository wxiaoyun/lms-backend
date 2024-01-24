package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SetupFrontend(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/assets") {
			return c.Next()
		}

		if strings.HasPrefix(c.Path(), "/file") {
			return c.Next()
		}

		if strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}

		// For any other unprompted routes, redirect to index.html
		return c.SendFile("./frontend/index.html")
	})
}
