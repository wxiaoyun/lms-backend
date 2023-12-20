package bookview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/loanview"
	"time"
)

type BookLoanView struct {
	BaseView
	Loan loanview.View `json:"loan"`
}

func ToBookLoanView(book *model.Book, loan *model.Loan) *BookLoanView {
	return &BookLoanView{
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
		Loan: *loanview.ToView(loan),
	}
}
