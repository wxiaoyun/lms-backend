package userhandler

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/bookmark"
	"lms-backend/internal/dataaccess/fine"
	"lms-backend/internal/dataaccess/loan"
	"lms-backend/internal/dataaccess/reservation"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/session"
	"lms-backend/internal/view/userview"

	"github.com/gofiber/fiber/v2"
)

func HandleGetCurrentUser(c *fiber.Ctx) error {
	sess, err := session.Store.Get(c)
	if err != nil {
		return err
	}

	token := sess.Get(session.CookieKey)
	userID, ok := token.(uint)
	if !ok || userID == 0 {
		return c.JSON(api.Response{
			Data: userview.ToGuestView(),
			Messages: api.Messages(
				api.SuccessMessage("Welcome guest!"),
			),
		})
	}

	id := int64(userID)

	db := database.GetDB()

	usr, err := user.Read(db, id)
	if err != nil {
		return err
	}

	abilites, err := user.GetAbilities(db, id)
	if err != nil {
		return err
	}

	bookmarks, err := bookmark.ListByUserID(db, id)
	if err != nil {
		return err
	}

	loans, err := loan.ListBorrowedLoanByUserID(db, id)
	if err != nil {
		return err
	}

	reservations, err := reservation.ListPendingReservationByUserID(db, id)
	if err != nil {
		return err
	}

	fines, err := fine.ListOutstandingFineByUserID(db, id)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: userview.ToCurrentUserView(usr, abilites, bookmarks, loans, reservations, fines),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf("Welcome back, %s!", usr.Username)),
		),
	})
}
