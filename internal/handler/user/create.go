package userhandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/params/userparams"
	"lms-backend/internal/view/userview"
)

func HandleCreate(c *fiber.Ctx) error {
	var params userparams.CreateParams
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return err
	}

	usr := params.ToModel()
	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("create new user %s", usr.Username),
	)
	defer func() { rollBackOrCommit(err) }()

	usr, err = user.Create(tx, usr)
	if err != nil {
		return err
	}

	abilities, err := user.GetAbilities(tx, int64(usr.ID))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: userview.ToView(usr, abilities...),
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"User %s created successfully", usr.Username,
			))),
	})
}
