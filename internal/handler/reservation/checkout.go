package reservationhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/bookcopy"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/reservationpolicy"
	"lms-backend/internal/session"
	"lms-backend/internal/view/reservationview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	checkoutBookAction = "checkout book"
)

func HandleCheckout(c *fiber.Ctx) error {
	param := c.Params("reservation_id")
	resID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid reservation id.", param))
	}

	err = policy.Authorize(c, checkoutBookAction, reservationpolicy.CheckoutPolicy(resID))
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
		c, fmt.Sprintf("%s checking out reservation id - \"%d\"", username, resID),
	)
	defer func() { rollBackOrCommit(err) }()

	res, err := bookcopy.CheckOutCopy(tx, userID, resID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: reservationview.ToDetailedView(res),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"%s has checked out reservation id - \"%d\".", username, resID,
			))),
	})
}
