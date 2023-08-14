package userhandler

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/session"
	"lms-backend/internal/view/userview"

	"github.com/gofiber/fiber/v2"
)

func HandleGetCurrentUser(c *fiber.Ctx) error {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	db := database.GetDB()

	usr, err := user.Read(db, userID)
	if err != nil {
		return err
	}

	abilites, err := user.GetAbilities(db, userID)
	if err != nil {
		return err
	}

	view := userview.ToView(usr, abilites)

	return c.JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf("Welcome back, %s!", usr.Username)),
		),
	})
}
