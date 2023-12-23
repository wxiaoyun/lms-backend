package fineview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/userview"
	"time"

	"github.com/ForAeons/ternary"
)

type LoanView struct {
	ID         int64      `json:"id,omitempty"`
	UserID     int64      `json:"user_id"`
	BookID     int64      `json:"book_id"`
	Status     string     `json:"status"`
	BorrowDate *time.Time `json:"borrow_date"`
	DueDate    *time.Time `json:"due_date"`
	ReturnDate *time.Time `json:"return_date"`
}

type BookView struct {
	ID              uint   `json:"id,omitempty"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	Publisher       string `json:"publisher"`
	PublicationDate string `json:"publication_date"`
	Genre           string `json:"genre"`
	Language        string `json:"language"`
}

type DetailedView struct {
	BaseView
	Book BookView      `json:"book"`
	Loan LoanView      `json:"loan"`
	User userview.View `json:"user"`
}

func ToDetailedView(fine *model.Fine) *DetailedView {
	return &DetailedView{
		BaseView: *ToBaseView(fine),
		User:     *userview.ToView(fine.User),
		Loan: LoanView{
			ID:         int64(fine.Loan.ID),
			UserID:     int64(fine.Loan.UserID),
			BookID:     int64(fine.Loan.BookID),
			Status:     fine.Loan.Status,
			BorrowDate: &fine.Loan.BorrowDate,
			DueDate:    &fine.Loan.DueDate,
			ReturnDate: ternary.If[*time.Time](fine.Loan.ReturnDate.Valid).
				Then(&fine.Loan.ReturnDate.Time).
				Else(nil),
		},
		Book: BookView{
			ID:              fine.Loan.Book.ID,
			Title:           fine.Loan.Book.Title,
			Author:          fine.Loan.Book.Author,
			ISBN:            fine.Loan.Book.ISBN,
			Publisher:       fine.Loan.Book.Publisher,
			PublicationDate: fine.Loan.Book.PublicationDate.Format(time.RFC3339),
			Genre:           fine.Loan.Book.Genre,
			Language:        fine.Loan.Book.Language,
		},
	}
}
