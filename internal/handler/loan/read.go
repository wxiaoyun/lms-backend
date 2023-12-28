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

func HandleRead(c *fiber.Ctx) error {
	err := policy.Authorize(c, readLoanAction, loanpolicy.ReadPolicy())
	if err != nil {
		return err
	}

	// userID, err := session.GetLoginSession(c)
	// if err != nil {
	// 	return err
	// }

	param := c.Params("loan_id")
	loanID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid loan id.", param))
	}

	db := database.GetDB()

	loanModel, err := loan.ReadDetailed(db, loanID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: loanview.ToDetailedView(loanModel),
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"Loan %d retrieved", loanID,
			))),
	})
}
