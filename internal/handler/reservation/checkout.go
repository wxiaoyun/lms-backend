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
// @Produce application/json
// @Success 200 {object} api.SwgResponse[reservationview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book/{book_id}/reservation/{reservation_id}/checkout [patch]
func HandleCheckout(c *fiber.Ctx) error {
	param := c.Params("id")
	bookID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	}
	param2 := c.Params("reservation_id")
	resID, err := strconv.ParseInt(param2, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid reservation id.", param2))
	}

	err = policy.Authorize(c, checkoutBookAction, reservationpolicy.CheckoutPolicy(resID, bookID))
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

	bookTitle, err := book.GetBookTitle(db, bookID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s checking out reservation for \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	res, err := book.CheckOutReservation(tx, userID, bookID, resID)
	if err != nil {
		return err
	}

	if res.BookID != uint(bookID) {
		err = externalerrors.BadRequest(fmt.Sprintf(
			"Reservation with id %d is not for book with id %d.", resID, bookID,
		))
		return err
	}

	view := reservationview.ToView(res)

	return c.JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"%s has checked out \"%s\".", username, bookTitle,
			))),
	})
}
