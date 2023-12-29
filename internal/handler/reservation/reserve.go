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
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	reserveBookAction = "reserve book"
)

func HandleReserve(c *fiber.Ctx) error {
	err := policy.Authorize(c, reserveBookAction, reservationpolicy.ReservePolicy())
	if err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	param := c.Params("book_id") // This is actually book copy id
	copyID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book copy id.", param))
	}

	db := database.GetDB()

	username, err := user.GetUserName(db, userID)
	if err != nil {
		return err
	}

	bookTitle, err := bookcopy.GetBookTitle(db, copyID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s reserving \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	res, err := bookcopy.ReserveCopy(tx, userID, copyID)
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
