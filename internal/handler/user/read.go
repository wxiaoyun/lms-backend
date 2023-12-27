package userhandler

import (
	"fmt"
	"lms-backend/internal/api"
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
	readUserAction = "update user"
)

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

	abilities, err := user.GetAbilities(db, userID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: userview.ToView(usr, abilities...),
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"User %s retrieved successfully", usr.Username,
			))),
	})
}
