package health

import (
	"lms-backend/internal/api"

	"github.com/gofiber/fiber/v2"
)

func HandleHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data: true,
		Messages: api.Messages(
			api.SilentMessage("server is running"),
		),
	})
}
