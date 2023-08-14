package reservationhandler

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/reservation"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/reservationpolicy"
	"lms-backend/internal/view/reservationview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	readReservationAction = "read reservation"
)

func HandleRead(c *fiber.Ctx) error {
	err := policy.Authorize(c, readReservationAction, reservationpolicy.DeletePolicy())
	if err != nil {
		return err
	}

	// userID, err := session.GetLoginSession(c)
	// if err != nil {
	// 	return err
	// }

	// param := c.Params("id")
	// bookID, err := strconv.ParseInt(param, 10, 64)
	// if err != nil {
	// 	return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	// }
	param2 := c.Params("reservation_id")
	resID, err := strconv.ParseInt(param2, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid reservation id.", param2))
	}

	db := database.GetDB()

	res, err := reservation.Read(db, resID)
	if err != nil {
		return err
	}

	view := reservationview.ToView(res)

	return c.JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"Reservation %d retrieved.", resID,
			))),
	})
}
