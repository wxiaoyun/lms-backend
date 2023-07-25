package userhandler

import (
	"fmt"
	"technical-test/internal/api"
	"technical-test/internal/dataaccess/user"
	"technical-test/internal/database"
	"technical-test/internal/session"
	"technical-test/internal/view/userview"

	"github.com/gofiber/fiber/v2"
)

func HandleGetCurrentUser(c *fiber.Ctx) error {
	sess, err := session.GetLoginSession(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(api.Response{
			Messages: []api.Message{api.InfoMessage("User is not logged in")},
		})
	}

	db := database.GetDB()

	user1, err := user.ReadByEmail(db, sess.Email)
	if err != nil {
		return err
	}

	view := userview.ToView(user1)

	return c.JSON(api.Response{
		Data: view,
		Messages: []api.Message{
			api.SuccessMessage(fmt.Sprintf("Welcome back, user %s!", user1.Email)),
		},
	})
}
