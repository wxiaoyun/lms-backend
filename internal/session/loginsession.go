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

func GetLoginSession(c *fiber.Ctx) (int64, error) {
	sess, err := Store.Get(c)
	if err != nil {
		return 0, err
	}

	token := sess.Get(CookieKey)
	if token == nil {
		err := c.Status(fiber.StatusUnauthorized).JSON(api.Response{
			Messages: []api.Message{api.InfoMessage("User is not logged in")},
		})
		if err != nil {
			return 0, err
		}
		return 0, fiber.NewError(fiber.StatusUnauthorized, "User is not logged in")
	}

	userID, ok := token.(uint)
	if !ok {
		return 0, fiber.NewError(fiber.StatusInternalServerError, "Erro casting session to integer")
	}

	if userID == 0 {
		err := c.Status(fiber.StatusUnauthorized).JSON(api.Response{
			Messages: []api.Message{api.InfoMessage("User is not logged in")},
		})
		if err != nil {
			return 0, err
		}
		return 0, fiber.NewError(fiber.StatusUnauthorized, "User is not logged in")
	}

	return int64(userID), nil
}
