package fineview

import (
	"lms-backend/internal/model"
)

type View struct {
	ID     int64   `json:"id,omitempty"`
	UserID int64   `json:"user_id"`
	LoanID int64   `json:"loan_id"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

func ToView(fine *model.Fine) *View {
	return &View{
		ID:     int64(fine.ID),
		UserID: int64(fine.UserID),
		LoanID: int64(fine.LoanID),
		Status: fine.Status,
		Amount: fine.Amount,
	}
}

func ToViews(fines []model.Fine) []View {
	views := make([]View, 0, len(fines))
	for _, fine := range fines {
		//nolint
		views = append(views, *ToView(&fine))
	}
	return views
}
