package health

import (
	"lms-backend/internal/api"

	"github.com/gofiber/fiber/v2"
)

// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /api/v1/health [get]
func HandleHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(api.Response{
		Messages: []api.Message{
			api.SilentMessage("server is running"),
		},
	})
}
