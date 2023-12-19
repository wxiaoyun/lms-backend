package reservationpolicy

import (
	"fmt"
	"lms-backend/internal/database"
	"lms-backend/internal/model"
	"lms-backend/internal/policy"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

type ReservationBelongsToUser struct {
	ReservationID int64
}

func AllowIfReservationBelongsToUser(reservationID int64) *ReservationBelongsToUser {
	return &ReservationBelongsToUser{
		ReservationID: reservationID,
	}
}

func (p *ReservationBelongsToUser) Validate(c *fiber.Ctx) (policy.Decision, error) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return policy.Deny, err
	}

	db := database.GetDB()

	var exists int64
	result := db.Model(&model.Reservation{}).
		Where("id = ? AND user_id = ?", p.ReservationID, userID).
		Count(&exists)
	if result.Error != nil {
		return policy.Deny, result.Error
	}

	if exists == 0 {
		return policy.Deny, nil
	}

	return policy.Allow, nil
}

func (p *ReservationBelongsToUser) Reason() string {
	return fmt.Sprintf("Reservation with ID %d does not belong to you.", p.ReservationID)
}
