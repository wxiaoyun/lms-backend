package bookview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type View struct {
	sharedview.BookView
}

func ToView(book *model.Book) *View {
	return &View{
		BookView: *sharedview.ToBookView(book),
	}
}
