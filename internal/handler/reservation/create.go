package reservationhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/params/sharedparams"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/reservationpolicy"
	"lms-backend/internal/view/reservationview"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	createReservationAction = "create reservation"
)

// @Summary Admin reservations a book on behalf of a user
// @Description Admin reservations a book on behalf of a user
// @Tags loan
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[reservationview.DetailedView]
// @Failure 400 {object} api.SwgErrResponse
// @Router /v1/reservation/ [post]
func HandleCreate(c *fiber.Ctx) error {
	err := policy.Authorize(c, createReservationAction, reservationpolicy.CreatePolicy())
	if err != nil {
		return err
	}

	var params sharedparams.UserBookParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return err
	}

	db := database.GetDB()

	username, err := user.GetUserName(db, params.UserID)
	if err != nil {
		return err
	}

	bookTitle, err := book.GetBookTitle(db, params.BookID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s loaning \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	res, err := book.ReserveBook(tx, params.UserID, params.BookID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: reservationview.ToDetailedView(res),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"\"%s\" has been reserved until %s.", bookTitle,
				res.ReservationDate.Format(time.RFC3339),
			))),
	})
}
