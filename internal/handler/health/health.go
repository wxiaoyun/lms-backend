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
// @Router /v1/health [get]
func HandleHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data: true,
		Messages: api.Messages(
			api.SilentMessage("server is running"),
		),
	})
}
