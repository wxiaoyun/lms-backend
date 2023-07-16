package auth

import "github.com/gofiber/fiber/v2"

func HandleSignIn(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Sign In")
}
