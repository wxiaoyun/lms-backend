package bookhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"
	"lms-backend/internal/view/bookview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Delete a book
// @Description deletes a book from the library
// @Tags book
// @Accept */*
// @Param book_id path int true "Book ID to delete"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[bookview.BaseView]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book/{book_id} [delete]
func HandleDelete(c *fiber.Ctx) error {
	err := policy.Authorize(c, createBookAction, bookpolicy.DeletePolicy())
	if err != nil {
		return err
	}

	param := c.Params("book_id")
	bookID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("Deleting a book in library - ID: %d", bookID),
	)
	defer func() { rollBackOrCommit(err) }()

	bookModel, err := book.Delete(tx, bookID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: bookview.ToView(bookModel),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"\"%s\" removed from library.", bookModel.Title,
			))),
	})
}
