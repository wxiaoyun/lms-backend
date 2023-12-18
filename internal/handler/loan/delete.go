package loanhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/loan"
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
	deleteLoanAction = "delete loan"
)

// @Summary Delete a loan
// @Description Deletes an existing loan in the library
// @Tags loan
// @Accept */*
// @Param loan_id path int true "loan ID to delete"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[loanview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/loan/{loan_id}/ [delete]
func HandleDelete(c *fiber.Ctx) error {
	err := policy.Authorize(c, deleteLoanAction, loanpolicy.DeletePolicy())
	if err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	param2 := c.Params("loan_id")
	loanID, err := strconv.ParseInt(param2, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid loan id.", param2))
	}

	db := database.GetDB()

	username, err := user.GetUserName(db, userID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s deleting loan of id - %d ", username, loanID),
	)
	defer func() { rollBackOrCommit(err) }()

	ln, err := loan.Delete(tx, loanID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: loanview.ToView(ln),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Loan id - \"%d\" has been deleted", loanID,
			))),
	})
}
