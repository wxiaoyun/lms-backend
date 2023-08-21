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
	renewLoanAction = "renew loan"
)

// @Summary Renew a loan
// @Description Renews a loan from the library
// @Tags loan
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[loanview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book/{book_id}/loan/{loan_id}/renew [patch]
func HandleRenew(c *fiber.Ctx) error {
	param := c.Params("book_id")
	bookID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	}
	param2 := c.Params("loan_id")
	loanID, err := strconv.ParseInt(param2, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid loan id.", param2))
	}

	err = policy.Authorize(c, renewLoanAction, loanpolicy.RenewPolicy(loanID, bookID))
	if err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
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
		c, fmt.Sprintf("%s renewing loan for \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	ln, err := book.RenewLoan(tx, loanID)
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
		Data: loanview.ToView(ln),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Loan for \"%s\" has been extended to %s.", bookTitle,
				ln.DueDate.Format(time.RFC3339),
			))),
	})
}
