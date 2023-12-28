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
	settleFineAction = "settle Fine"
)

func HandleSettle(c *fiber.Ctx) error {
	param3 := c.Params("fine_id")
	fineID, err := strconv.ParseInt(param3, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid fine id.", param3))
	}

	err = policy.Authorize(c, settleFineAction, finepolicy.SettlePolicy(fineID))
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
		c, fmt.Sprintf("%s settling fine id: \"%d\"", username, fineID),
	)
	defer func() { rollBackOrCommit(err) }()

	fn, err := fine.Settle(tx, fineID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: fineview.ToDetailedView(fn),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Fine id - \"%d\" is settled.", fineID,
			))),
	})
}
