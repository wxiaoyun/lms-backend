package finehandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
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
// @Param fine_id path int true "fine ID to delete"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[fineview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /v1/fine/{fine_id} [delete]
func HandleDelete(c *fiber.Ctx) error {
	err := policy.Authorize(c, deleteFineAction, finepolicy.DeletePolicy())
	if err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	param := c.Params("fine_id")
	fineID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid fine id.", param))
	}

	db := database.GetDB()
	username, err := user.GetUserName(db, userID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s deleting fine of id - %d ", username, fineID),
	)
	defer func() { rollBackOrCommit(err) }()

	fn, err := fine.Delete(tx, fineID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: fineview.ToDetailedView(fn),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Fine id - \"%d\" has been deleted", fineID,
			))),
	})
}
