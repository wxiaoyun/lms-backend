package health

import (
	"lms-backend/internal/api"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func HandleHealth(c *fiber.Ctx) error {
	csrfToken, ok := c.Locals(middleware.CSRFContextKey).(string)
	if !ok {
		csrfToken = ""
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data: csrfToken,
		Messages: api.Messages(
			api.SilentMessage("server is running"),
		),
	})
}
