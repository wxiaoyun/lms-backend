package reservationhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
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
	deleteReservationAction = "delete reservation"
)

// @Summary Delete a reservation
// @Description Deletes a reservation for a book
// @Tags reservation
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[reservationview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book/{book_id}/reservation/{reservation_id}/ [delete]
func HandleDelete(c *fiber.Ctx) error {
	err := policy.Authorize(c, deleteReservationAction, reservationpolicy.DeletePolicy())
	if err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

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
		c, fmt.Sprintf("%s deleting reservation ID - %d belonging to \"%s\"", username, resID, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	res, err := reservation.Delete(tx, resID)
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
				"Reservation for \"%s\" is deleted.", bookTitle,
			))),
	})
}
