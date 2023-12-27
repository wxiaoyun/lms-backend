package bookhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/params/bookparams"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"
	"lms-backend/internal/view/bookview"

	"github.com/gofiber/fiber/v2"
)

const (
	createBookAction = "create book"
)

func HandleCreate(c *fiber.Ctx) error {
	err := policy.Authorize(c, createBookAction, bookpolicy.CreatePolicy())
	if err != nil {
		return err
	}

	var bookParams bookparams.CreateParams
	if err := c.BodyParser(&bookParams); err != nil {
		return err
	}

	if err := bookParams.Validate(); err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("Adding a new book to library: %s.", bookParams.Title),
	)
	defer func() { rollBackOrCommit(err) }()

	bookModel := bookParams.ToModel()
	bookModel, err = book.Create(tx, bookModel)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: bookview.ToView(bookModel),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"\"%s\" added to library.", bookModel.Title,
			))),
	})
}
