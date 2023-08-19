package health

import (
	"lms-backend/internal/api"

	"github.com/gofiber/fiber/v2"
)

// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgMsgResponse
// @Router /api/v1/health [get]
func HandleHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(api.Response{
		Messages: api.Messages(
			api.SilentMessage("server is running"),
		),
	})
}
