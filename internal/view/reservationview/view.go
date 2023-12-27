package reservationview

import (
	"lms-backend/internal/model"
	"time"
)

type BaseView struct {
	ID              int64     `json:"id,omitempty"`
	UserID          int64     `json:"user_id"`
	BookID          int64     `json:"book_id"`
	Status          string    `json:"status"`
	ReservationDate time.Time `json:"reservation_date"`
}

func ToView(reservation *model.Reservation) *BaseView {
	return &BaseView{
		ID:              int64(reservation.ID),
		UserID:          int64(reservation.UserID),
		BookID:          int64(reservation.BookCopyID),
		Status:          reservation.Status,
		ReservationDate: reservation.ReservationDate,
	}
}

func ToViews(reservations []model.Reservation) []BaseView {
	views := make([]BaseView, 0, len(reservations))
	for _, reservation := range reservations {
		//nolint
		views = append(views, *ToView(&reservation))
	}
	return views
}
