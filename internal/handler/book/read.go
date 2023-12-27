package bookhandler

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"
	"lms-backend/internal/view/bookview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	readBookAction = "read book"
)

func HandleRead(c *fiber.Ctx) error {
	err := policy.Authorize(c, readBookAction, bookpolicy.ReadPolicy())
	if err != nil {
		return err
	}

	param := c.Params("book_id")
	bookID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	}

	db := database.GetDB()

	bookModel, err := book.ReadWithCopies(db, bookID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: bookview.ToDetailedView(bookModel),
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"\"%s\" retrieved.", bookModel.Title,
			))),
	})
}
