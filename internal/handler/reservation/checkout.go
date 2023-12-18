package reservationhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
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

// @Summary Checkout a book
// @Description Checks out a book for a given reservation
// @Tags reservation
// @Accept */*
// @Param reservation_id path int true "reservation ID to checkout"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[reservationview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/reservation/{reservation_id}/checkout [patch]
func HandleCheckout(c *fiber.Ctx) error {
	param2 := c.Params("reservation_id")
	resID, err := strconv.ParseInt(param2, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid reservation id.", param2))
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

	res, err := book.CheckOutReservation(tx, userID, resID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: reservationview.ToView(res),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"%s has checked out reservation id - \"%d\".", username, resID,
			))),
	})
}
