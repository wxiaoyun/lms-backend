package userhandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/model"
	"lms-backend/internal/params/userparams"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/userpolicy"
	"lms-backend/internal/view/userview"
	"lms-backend/pkg/error/externalerrors"
)

const (
	updateUserAction = "update user"
)

// @Summary Update an existing user
// @Description Updates an existing user in the system
// @Tags user
// @Accept application/json
// @Param createuserparam body userparams.UpdateParams true "User update request"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[userview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/user/{user_id} [patch]
func HandleUpdate(c *fiber.Ctx) error {
	param := c.Params("user_id")
	userID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid user id.", param))
	}

	err = policy.Authorize(c, updateUserAction, userpolicy.UpdatePolicy(userID))
	if err != nil {
		return err
	}

	var params userparams.UpdateParams
	err = c.BodyParser(&params)
	if err != nil {
		return err
	}

	if err := params.Validate(userID); err != nil {
		return err
	}

	usr := params.ToModel()
	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("updating user %s", usr.Username),
	)
	defer func() { rollBackOrCommit(err) }()

	usr, err = user.Update(tx, usr)
	if err != nil {
		return err
	}

	view := userview.ToView(usr, []model.Ability{})

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"User %s updated successfully", usr.Username,
			))),
	})
}
