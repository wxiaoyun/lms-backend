package session

import (
	"lms-backend/internal/api"

	"github.com/gofiber/fiber/v2"
)

type LoginSession struct {
	UserID         uint
	Email          string
	IsMasquerading bool
}

func HasSession(c *fiber.Ctx) bool {
	token := c.Locals(UserIDKey)
	if token == nil {
		return false
	}

	userID, ok := token.(uint)
	if !ok {
		return false
	}

	if userID == 0 {
		return false
	}

	return true
}

func GetLoginSession(c *fiber.Ctx) (int64, error) {
	token := c.Locals(UserIDKey)
	if token == nil {
		//nolint
		c.Status(fiber.StatusUnauthorized).JSON(api.Response{
			Messages: []api.Message{api.InfoMessage("User is not logged in")},
		})
		return 0, fiber.NewError(fiber.StatusUnauthorized, "User is not logged in")
	}

	userID, ok := token.(uint)
	if !ok {
		return 0, fiber.NewError(fiber.StatusInternalServerError, "Erro casting session to integer")
	}

	if userID == 0 {
		//nolint
		c.Status(fiber.StatusUnauthorized).JSON(api.Response{
			Messages: []api.Message{api.InfoMessage("User is not logged in")},
		})
		return 0, fiber.NewError(fiber.StatusUnauthorized, "User is not logged in")
	}

	return int64(userID), nil
}
