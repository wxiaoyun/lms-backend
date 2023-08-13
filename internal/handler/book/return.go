package bookhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"
	"lms-backend/internal/session"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandleReturn(c *fiber.Ctx) error {
	err := policy.Authorize(c, readBookAction, bookpolicy.RenewPolicy())
	if err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	param := c.Params("id")
	bookID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	}

	db := database.GetDB()

	username, err := user.GetUserName(db, userID)
	if err != nil {
		return err
	}

	bookTitle, err := book.GetBookTitle(db, bookID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, db, fmt.Sprintf("%s returning \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	bookModel, _, err := book.ReturnBook(tx, userID, bookID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Messages: []api.Message{
			api.SuccessMessage(fmt.Sprintf(
				"Book \"%s\" has been returned.", bookModel.Title,
			))},
	})
}
