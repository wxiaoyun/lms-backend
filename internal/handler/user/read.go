package userhandler

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/model"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/userpolicy"
	"lms-backend/internal/view/userview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	readUserAction = "update user"
)

// @Summary Read an existing user
// @Description Updates an existing user in the system
// @Tags user
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[userview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/user/{user_id} [get]
func HandleRead(c *fiber.Ctx) error {
	param := c.Params("user_id")
	userID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid user id.", param))
	}

	err = policy.Authorize(c, readUserAction, userpolicy.ReadPolicy(userID))
	if err != nil {
		return err
	}

	db := database.GetDB()

	usr, err := user.Read(db, userID)
	if err != nil {
		return err
	}

	view := userview.ToView(usr, []model.Ability{})

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"User %s retrieved successfully", usr.Username,
			))),
	})
}
