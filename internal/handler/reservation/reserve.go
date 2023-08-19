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
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	reserveBookAction = "reserve book"
)

// @Summary Create a reservation
// @Description Creates a reservation for a book
// @Tags reservation
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[reservationview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book/{book_id}/reservation/ [post]
func HandleReserve(c *fiber.Ctx) error {
	err := policy.Authorize(c, reserveBookAction, reservationpolicy.ReservePolicy())
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
		c, db, fmt.Sprintf("%s reserving \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	res, err := book.ReserveBook(tx, userID, bookID)
	if err != nil {
		return err
	}

	view := reservationview.ToView(res)

	return c.JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"\"%s\" has been reserved until %s.", bookTitle,
				res.ReservationDate.Format(time.RFC3339),
			))),
	})
}
