package loanhistoryview

import (
	"lms-backend/internal/model"
)

type View struct {
	ID     int64  `json:"id,omitempty"`
	LoanID int64  `json:"loan_id"`
	Action string `json:"action"`
}

func ToView(loan *model.LoanHistory) *View {
	return &View{
		ID:     int64(loan.ID),
		LoanID: int64(loan.LoanID),
		Action: loan.Action,
	}
}

func ToViews(loans []model.LoanHistory) []View {
	views := make([]View, 0, len(loans))
	for _, loan := range loans {
		//nolint
		views = append(views, *ToView(&loan))
	}
	return views
}
