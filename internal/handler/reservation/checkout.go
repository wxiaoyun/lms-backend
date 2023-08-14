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
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	checkoutBookAction = "checkout book"
)

func HandleCheckout(c *fiber.Ctx) error {
	err := policy.Authorize(c, checkoutBookAction, reservationpolicy.CheckoutPolicy())
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
		c, db, fmt.Sprintf("%s checking out reservation for \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	_, err = book.CheckOutReservation(tx, userID, bookID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"%s has checked out \"%s\".", username, bookTitle,
			))),
	})
}
