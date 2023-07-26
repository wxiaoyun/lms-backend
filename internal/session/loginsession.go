package session

import (
	"fmt"
	"technical-test/internal/api"

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

		c.Status(fiber.StatusUnauthorized).JSON(api.Response{
			Messages: []api.Message{api.InfoMessage("User is not logged in")},
		})
		return 0, fiber.NewError(fiber.StatusUnauthorized, "User is not logged in")
	}
	fmt.Println("Mdware1: ", token)

	userID, ok := token.(uint)
	if !ok {
		return 0, fiber.NewError(fiber.StatusInternalServerError, "Erro casting session to integer")
	}
	fmt.Println("Mdware2: ", userID)

	if userID == 0 {
		//nolint:errcheck // this always return nil error
		c.Status(fiber.StatusUnauthorized).JSON(api.Response{
			Messages: []api.Message{api.InfoMessage("User is not logged in")},
		})
		return 0, fiber.NewError(fiber.StatusUnauthorized, "User is not logged in")
	}

	fmt.Println("Mdware3: ", userID)

	return int64(userID), nil
}
