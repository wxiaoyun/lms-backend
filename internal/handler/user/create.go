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

// @Summary Create a new user
// @Description Creates a new user in the system
// @Tags auth
// @Accept application/json
// @Param createuserparam body userparams.CreateParams true "User creation request"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[userview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/auth/signup [post]
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
