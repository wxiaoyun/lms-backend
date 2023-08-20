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

// @Summary Update a book
// @Description Updates an existing book in the library
// @Tags book
// @Accept application/json
// @Param book body bookparams.UpdateParams true "Book update request"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[bookview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book/ [patch]
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

	view := bookview.ToView(bookModel)

	return c.JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"\"%s\" modified successfully.", bookModel.Title,
			))),
	})
}
