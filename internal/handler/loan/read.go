package loanhandler

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/loan"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/loanpolicy"
	"lms-backend/internal/view/loanview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	readLoanAction = "read loan"
)

// @Summary Read a loan
// @Description Retrieves a loan from the library
// @Tags loan
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[loanview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book/{book_id}/loan/{loan_id}/ [get]
func HandleRead(c *fiber.Ctx) error {
	err := policy.Authorize(c, readLoanAction, loanpolicy.ReadPolicy())
	if err != nil {
		return err
	}

	// userID, err := session.GetLoginSession(c)
	// if err != nil {
	// 	return err
	// }

	// param := c.Params("id")
	// bookID, err := strconv.ParseInt(param, 10, 64)
	// if err != nil {
	// 	return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	// }
	param2 := c.Params("loan_id")
	loanID, err := strconv.ParseInt(param2, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid loan id.", param2))
	}

	db := database.GetDB()

	loanModel, err := loan.ReadDetailed(db, loanID)
	if err != nil {
		return err
	}

	view := loanview.ToView(loanModel)

	return c.JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"Loan %d retrieved", loanID,
			))),
	})
}
