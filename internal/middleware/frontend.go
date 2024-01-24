package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SetupWebApp(app *fiber.App) {
	// For the favicon, respond with icon.svg
	app.Get("/icon.svg", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/icon.svg")
	})

	// For the root path, respond with index.html
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/index.html")
	})

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
