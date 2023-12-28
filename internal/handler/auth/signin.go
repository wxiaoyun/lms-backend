package auth

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/bookmark"
	"lms-backend/internal/dataaccess/fine"
	"lms-backend/internal/dataaccess/loan"
	"lms-backend/internal/dataaccess/reservation"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/params/userparams"
	"lms-backend/internal/session"
	"lms-backend/internal/view/userview"

	"github.com/gofiber/fiber/v2"
)

func HandleSignIn(c *fiber.Ctx) error {
	// if handler, ok := c.Locals(csrf.ConfigDefault.HandlerContextKey).(*csrf.CSRFHandler); ok {
	// 	if err := handler.DeleteToken(c); err != nil {
	// 		return err
	// 	}
	// }

	var params userparams.SignInParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return err
	}

	usr := params.ToModel()
	db := database.GetDB()
	usr, err := user.Login(db, usr)
	if err != nil {
		return err
	}

	sess, err := session.Store.Get(c)
	if err != nil {
		return err
	}

	err = sess.Regenerate()
	if err != nil {
		return err
	}

	sess.Set(session.CookieKey, usr.ID)
	err = sess.Save()
	if err != nil {
		return err
	}

	id := int64(usr.ID)

	abilities, err := user.GetAbilities(db, int64(usr.ID))
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

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data: userview.ToLoginView(usr, abilities, bookmarks, loans, reservations, fines),
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"%s is logged in successfully", usr.Username,
			))),
	})
}
