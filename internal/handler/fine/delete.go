package finehandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/dataaccess/fine"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/finepolicy"
	"lms-backend/internal/session"
	"lms-backend/internal/view/fineview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	deleteFineAction = "delete fine"
)

// @Summary Delete fine
// @Description deletes a fine belonging to a loan
// @Tags fine
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[fineview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book/{book_id}/loan/{loan_id}/fine/{fine_id} [delete]
func HandleDelete(c *fiber.Ctx) error {
	err := policy.Authorize(c, deleteFineAction, finepolicy.DeletePolicy())
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
	param2 := c.Params("loan_id")
	loanID, err := strconv.ParseInt(param2, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid loan id.", param2))
	}
	param3 := c.Params("fine_id")
	fineID, err := strconv.ParseInt(param3, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid fine id.", param3))
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
		c, fmt.Sprintf("%s deleting fine of id - %d belonging to \"%s\"", username, fineID, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	fn, err := fine.Delete(tx, fineID)
	if err != nil {
		return err
	}

	if fn.LoanID != uint(loanID) {
		err = externalerrors.BadRequest(fmt.Sprintf(
			"Fine with id %d does not belong to loan with id %d.", fn.ID, loanID,
		))
		return err
	}

	view := fineview.ToView(fn)

	return c.JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Fine for \"%s\" has been deleted", bookTitle,
			))),
	})
}
