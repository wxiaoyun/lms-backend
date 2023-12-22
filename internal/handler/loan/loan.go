package loanhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/loanpolicy"
	"lms-backend/internal/session"
	"lms-backend/internal/view/loanview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	loanBookAction = "loan book"
)

// @Summary Loan a book
// @Description Loans a book from the library
// @Tags loan
// @Accept */*
// @Param book_id path int true "Book ID for loan"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[loanview.DetailedView]
// @Failure 400 {object} api.SwgErrResponse
// @Router /v1/book/{book_id}/loan/ [post]
func HandleLoan(c *fiber.Ctx) error {
	err := policy.Authorize(c, loanBookAction, loanpolicy.LoanPolicy())
	if err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	param := c.Params("book_id")
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
		c, fmt.Sprintf("%s loaning \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	ln, err := book.LoanBook(tx, userID, bookID)
	if err != nil {
		return err
	}

	if ln.BookID != uint(bookID) {
		err = externalerrors.BadRequest(fmt.Sprintf(
			"Loan with id %d does not belong to %s.", ln.ID, bookTitle,
		))
		return err
	}

	return c.JSON(api.Response{
		Data: loanview.ToDetailedView(ln),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"\"%s\" is loaned until %s.", bookTitle,
				ln.DueDate.Format(time.RFC3339),
			))),
	})
}
