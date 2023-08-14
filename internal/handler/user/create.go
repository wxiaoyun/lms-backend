package userhandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/database"
	"lms-backend/internal/model"
	"lms-backend/internal/params/userparams"
	"lms-backend/internal/view/userview"
)

// @Summary Create a user
// @Description create an instance of user in the database
// @Tags user
// @Accept */*
// @Produce application/json
// @Success 200 "OK"
// @Router /api/v1/auth/signup [post]
func HandleCreateUser(c *fiber.Ctx) error {
	var params userparams.CreateUserParams
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	err = params.Validate()
	if err != nil {
		return err
	}

	user := params.ToModel()
	db := database.GetDB()
	tx, rollBackOrCommit := audit.Begin(
		c, db, fmt.Sprintf("create new user %s", user.Username),
	)
	defer func() { rollBackOrCommit(err) }()

	err = user.Create(tx)
	if err != nil {
		return err
	}

	view := userview.ToView(user, []model.Ability{})

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"User %s created successfully", user.Username,
			))),
	})
}
