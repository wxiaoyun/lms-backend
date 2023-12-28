package reservationview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type WithBookView struct {
	View
	Book *sharedview.BookView `json:"book"`
}

func ToBookView(res *model.Reservation) *WithBookView {
	return &WithBookView{
		View: *ToView(res),
		Book: sharedview.ToBookView(res.BookCopy.Book),
	}
}
