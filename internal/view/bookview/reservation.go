package bookview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/reservationview"
	"time"
)

type BookReservationView struct {
	BaseView
	Reservation reservationview.View `json:"reservation"`
}

func ToBookReservationView(book *model.Book, reservation *model.Reservation) *BookReservationView {
	return &BookReservationView{
		BaseView: BaseView{
			ID:              book.ID,
			Title:           book.Title,
			Author:          book.Author,
			ISBN:            book.ISBN,
			Publisher:       book.Publisher,
			PublicationDate: book.PublicationDate.Format(time.RFC3339),
			Genre:           book.Genre,
			Language:        book.Language,
		},
		Reservation: *reservationview.ToView(reservation),
	}
}
