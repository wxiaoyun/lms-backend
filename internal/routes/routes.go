package routes

import "github.com/gofiber/fiber/v2"

func SetUpRoutes() error {
	app := fiber.New()

	app.Listen(":3000")

	return nil
}
