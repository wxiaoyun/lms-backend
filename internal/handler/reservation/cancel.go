package reservationhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
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

// @Summary Cancel a reservation
// @Description Cancels a reservation for a book
// @Tags reservation
// @Accept */*
// @Param reservation_id path int true "reservation ID to cancel"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[reservationview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/reservation/{reservation_id}/cancel [patch]
func HandleCancel(c *fiber.Ctx) error {
	param2 := c.Params("reservation_id")
	resID, err := strconv.ParseInt(param2, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid reservation id.", param2))
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

	res, err := reservation.FullfilReservation(tx, resID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: reservationview.ToView(res),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Reservation id \"%d\" is canceled.", resID,
			))),
	})
}
