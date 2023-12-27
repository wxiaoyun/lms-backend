package reservationview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type View struct {
	sharedview.ResView
}

func ToView(reservation *model.Reservation) *View {
	return &View{
		ResView: *sharedview.ToResView(reservation),
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
