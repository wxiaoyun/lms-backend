package reservationview

import (
	"lms-backend/internal/model"
	"time"
)

type View struct {
	ID              int64     `json:"id,omitempty"`
	UserID          int64     `json:"user_id"`
	BookID          int64     `json:"book_id"`
	Status          string    `json:"status"`
	ReservationDate time.Time `json:"reservation_date"`
}

func ToView(reservation *model.Reservation) *View {
	return &View{
		ID:              int64(reservation.ID),
		UserID:          int64(reservation.UserID),
		BookID:          int64(reservation.BookID),
		Status:          reservation.Status,
		ReservationDate: reservation.ReservationDate,
	}
}

func ToViews(reservations []model.Reservation) []View {
	views := make([]View, 0, len(reservations))
	for _, reservation := range reservations {
		//nolint
		views = append(views, *ToView(&reservation))
	}
	return views
}
