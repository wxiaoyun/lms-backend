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
	"time"

	"github.com/gofiber/fiber/v2"
)

func HandleRenew(c *fiber.Ctx) error {
	err := policy.Authorize(c, readBookAction, bookpolicy.ReturnPolicy())
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
		c, db, fmt.Sprintf("%s renewing loan for \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	bookModel, loanModel, err := book.RenewLoan(tx, userID, bookID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Messages: []api.Message{
			api.SuccessMessage(fmt.Sprintf(
				"Loan for \"%s\" has been extended to %s.", bookModel.Title,
				loanModel.DueDate.Format(time.RFC3339),
			))},
	})
}
