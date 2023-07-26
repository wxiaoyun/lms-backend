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
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	db := database.GetDB()

	user1, err := user.Read(db, userID)
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
