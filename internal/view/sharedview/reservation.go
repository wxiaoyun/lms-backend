package sharedview

import (
	"lms-backend/internal/model"
	"time"
)

type ResView struct {
	ID              int64     `json:"id,omitempty"`
	UserID          int64     `json:"user_id"`
	BookCopyID      int64     `json:"book_copy_id"`
	Status          string    `json:"status"`
	ReservationDate time.Time `json:"reservation_date"`
}

func ToResView(reservation *model.Reservation) *ResView {
	return &ResView{
		ID:              int64(reservation.ID),
		UserID:          int64(reservation.UserID),
		BookCopyID:      int64(reservation.BookCopyID),
		Status:          reservation.Status,
		ReservationDate: reservation.ReservationDate,
	}
}
