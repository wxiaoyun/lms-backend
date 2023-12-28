package sharedview

import (
	"lms-backend/internal/model"
)

type FineView struct {
	ID     int64   `json:"id,omitempty"`
	UserID int64   `json:"user_id"`
	LoanID int64   `json:"loan_id"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

func ToFineView(fine *model.Fine) *FineView {
	return &FineView{
		ID:     int64(fine.ID),
		UserID: int64(fine.UserID),
		LoanID: int64(fine.LoanID),
		Status: fine.Status,
		Amount: fine.Amount,
	}
}
