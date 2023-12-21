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

	"github.com/gofiber/fiber/v2"
)

const (
	returnBookAction = "return book"
)

// @Summary Return a book
// @Description Returns a book to the library
// @Tags loan
// @Accept */*
// @Param loan_id path int true "loan ID to return"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[loanview.DetailedView]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/loan/{loan_id}/return [patch]
func HandleReturn(c *fiber.Ctx) error {
	param2 := c.Params("loan_id")
	loanID, err := strconv.ParseInt(param2, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid loan id.", param2))
	}

	err = policy.Authorize(c, returnBookAction, loanpolicy.ReturnPolicy(loanID))
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

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s returning loan id - \"%d\"", username, loanID),
	)
	defer func() { rollBackOrCommit(err) }()

	ln, err := book.ReturnBook(tx, loanID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: loanview.ToDetailedView(ln),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Loan id - \"%d\" has been returned.", loanID,
			))),
	})
}
