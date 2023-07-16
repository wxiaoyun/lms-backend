package health

import "github.com/gofiber/fiber/v2"

func HandleHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("I'm alive!")
}
