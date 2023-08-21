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

// @Summary Retrieve current user
// @Description Retrieves the current user if logged in
// @Tags user
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[userview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/user/currentuser [get]
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

	return c.JSON(api.Response{
		Data: userview.ToView(usr, abilites...),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf("Welcome back, %s!", usr.Username)),
		),
	})
}
