package bookcopyhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/dataaccess/bookcopy"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"
	"lms-backend/internal/view/sharedview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	createBookcopyAction = "create book copy"
	countQueryKey        = "count"
)

func HandleCreate(c *fiber.Ctx) error {
	err := policy.Authorize(c, createBookcopyAction, bookpolicy.CreatePolicy())
	if err != nil {
		return err
	}

	param := c.Params("book_id")
	bookID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	}

	count := c.QueryInt(countQueryKey, 1)

	db := database.GetDB()
	title, err := book.GetBookTitle(db, bookID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("Creating %d copies of %s", count, title),
	)
	defer func() { rollBackOrCommit(err) }()

	copies, err := bookcopy.CreateMultiple(tx, bookID, int64(count))
	if err != nil {
		return err
	}

	view := []sharedview.BookCopyView{}
	for _, w := range copies {
		//nolint:gosec // loop does not modify struct
		view = append(view, *sharedview.ToBookCopyView(&w))
	}

	return c.JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf("%d copies of %s created successfully", count, title)),
		),
	})
}
