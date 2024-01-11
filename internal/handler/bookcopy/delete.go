package bookcopyhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/bookcopy"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"
	"lms-backend/internal/view/bookcopyview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	deleteBookcopyAction = "delete book copy"
)

func HandleDelete(c *fiber.Ctx) error {
	err := policy.Authorize(c, deleteBookcopyAction, bookpolicy.DeletePolicy())
	if err != nil {
		return err
	}

	param := c.Params("bookcopy_id")
	bookcopyID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("Deleting book copy %d", bookcopyID),
	)
	defer func() { rollBackOrCommit(err) }()

	bookCopy, err := bookcopy.Delete(tx, bookcopyID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data:     bookcopyview.ToDetailedView(bookCopy),
		Messages: api.Messages(api.SilentMessage("book copy deleted successfully")),
	})
}
