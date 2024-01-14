package reservationhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/bookcopy"
	"lms-backend/internal/dataaccess/reservation"
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
	cancelReservationAction = "cancel reservation"
)

func HandleCancel(c *fiber.Ctx) error {
	param := c.Params("reservation_id")
	resID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid reservation id.", param))
	}

	err = policy.Authorize(c, cancelReservationAction, reservationpolicy.CancelPolicy(resID))
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
		c, fmt.Sprintf("%s canceling reservation id - \"%d\"", username, resID),
	)
	defer func() { rollBackOrCommit(err) }()

	res, err := bookcopy.CancelReservationCopy(tx, resID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: reservationview.ToDetailedView(res),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Reservation id \"%d\" is canceled.", resID,
			))),
	})
}

func HandleCancelByBookcopy(c *fiber.Ctx) error {
	param := c.Params("bookcopy_id")
	bookcopyID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book copy id.", param))
	}

	db := database.GetDB()

	res, err := reservation.ReadReservedByBookCopyID(db, bookcopyID)
	if err != nil {
		return err
	}

	err = policy.Authorize(c, cancelReservationAction, reservationpolicy.CancelPolicy(int64(res.ID)))
	if err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	username, err := user.GetUserName(db, userID)
	if err != nil {
		return err
	}

	title, err := bookcopy.GetBookTitle(db, bookcopyID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s canceling reservation for \"%s\"", username, title),
	)
	defer func() { rollBackOrCommit(err) }()

	res, err = bookcopy.CancelReservationCopy(tx, int64(res.ID))
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: reservationview.ToDetailedView(res),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Reservation id \"%d\" is canceled.", int64(res.ID),
			))),
	})
}
