package userhandler

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/userpolicy"
	"lms-backend/internal/view/userview"

	"github.com/gofiber/fiber/v2"
)

func HandleAutoComplete(c *fiber.Ctx) error {
	err := policy.Authorize(c, readUserAction, userpolicy.ListPolicy())
	if err != nil {
		return err
	}

	value := c.Params("value")

	db := database.GetDB()

	users, err := user.AutoComplete(db, value)
	if err != nil {
		return err
	}

	views := make([]*userview.SimpleView, len(users))
	for i, usr := range users {
		//nolint:gosec // loop does not modify struct
		views[i] = userview.ToSimpleView(&usr)
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: views,
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Autocomplete for \"%s\"", value,
			))),
	})
}
