package routes

import "github.com/gofiber/fiber/v2"

func SetUpRoutes() error {
	app := fiber.New()

	return app.Listen(":3000")
}
