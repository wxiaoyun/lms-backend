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
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandleUpdate(c *fiber.Ctx) error {
	err := policy.Authorize(c, createBookAction, bookpolicy.UpdatePolicy())
	if err != nil {
		return err
	}

	param := c.Params("book_id")
	bookID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	}

	var bookParams bookparams.UpdateParams
	if err := c.BodyParser(&bookParams); err != nil {
		return err
	}

	if err := bookParams.Validate(bookID); err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("Updating existing book in library: %s.", bookParams.Title),
	)
	defer func() { rollBackOrCommit(err) }()

	bookModel := bookParams.ToModel()
	bookModel, err = book.Update(tx, bookModel)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: bookview.ToView(bookModel),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"\"%s\" modified successfully.", bookModel.Title,
			))),
	})
}
