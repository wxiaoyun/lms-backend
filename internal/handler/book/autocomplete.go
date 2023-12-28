package bookhandler

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"
	"lms-backend/internal/view/bookview"

	"github.com/gofiber/fiber/v2"
)

func HandleAutoComplete(c *fiber.Ctx) error {
	err := policy.Authorize(c, readBookAction, bookpolicy.ListPolicy())
	if err != nil {
		return err
	}

	value := c.Params("value")

	db := database.GetDB()

	books, err := book.AutoComplete(db, value)
	if err != nil {
		return err
	}

	views := make([]*bookview.SimpleView, len(books))
	for i, usr := range books {
		//nolint:gosec // loop does not modify struct
		views[i] = bookview.ToSimpleView(&usr)
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: views,
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Autocomplete for \"%s\"", value,
			))),
	})
}
