package loanhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/dataaccess/bookcopy"
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

func HandleCreate(c *fiber.Ctx) error {
	err := policy.Authorize(c, createLoanAction, loanpolicy.CreatePolicy())
	if err != nil {
		return err
	}

	var params sharedparams.UserBookcopyParams
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

	bookTitle, err := bookcopy.GetBookTitle(db, params.BookCopyID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s loaning \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	ln, err := bookcopy.LoanCopy(tx, params.UserID, params.BookCopyID)
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

func HandleCreateByBook(c *fiber.Ctx) error {
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

	ln, err := book.Loan(tx, params.UserID, params.BookID)
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
