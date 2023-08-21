package userhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/userpolicy"
	"lms-backend/internal/view/userview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	deleteUserAction = "delete user"
)

// @Summary Delete an existing user
// @Description Deletes an existing user in the system
// @Tags user
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[userview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/user/{user_id} [delete]
func HandleDelete(c *fiber.Ctx) error {
	param := c.Params("user_id")
	userID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid user id.", param))
	}

	err = policy.Authorize(c, deleteUserAction, userpolicy.DeletePolicy(userID))
	if err != nil {
		return err
	}

	db := database.GetDB()
	username, err := user.GetUserName(db, userID)

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("deleting user %s", username),
	)
	defer func() { rollBackOrCommit(err) }()

	usr, err := user.Delete(tx, userID)
	if err != nil {
		return err
	}

	abilities, err := user.GetAbilities(tx, userID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: userview.ToView(usr, abilities),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"User %s deleted successfully", usr.Username,
			))),
	})
}
