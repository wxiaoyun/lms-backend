package loanhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/params/sharedparams"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/loanpolicy"
	"lms-backend/internal/view/loanview"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	createLoanAction = "loan book"
)

// @Summary Admin loans a book on behalf of a user
// @Description Admin loans a book on behalf of a user
// @Tags loan
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[loanview.DetailedView]
// @Failure 400 {object} api.SwgErrResponse
// @Router /v1/loan/ [post]
func HandleCreate(c *fiber.Ctx) error {
	err := policy.Authorize(c, createLoanAction, loanpolicy.CreatePolicy())
	if err != nil {
		return err
	}

	var params sharedparams.UserBookParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return err
	}

	db := database.GetDB()

	username, err := user.GetUserName(db, params.UserID)
	if err != nil {
		return err
	}

	bookTitle, err := book.GetBookTitle(db, params.BookID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s loaning \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	ln, err := book.LoanBook(tx, params.UserID, params.BookID)
	if err != nil {
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
