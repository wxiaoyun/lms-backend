package bookview

import "lms-backend/internal/viewmodel"

type PopularView struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	LoanCount int64  `json:"loan_count"`
}

func ToPopularView(b *viewmodel.BookLoanCount) *PopularView {
	return &PopularView{
		ID:        b.ID,
		Title:     b.Title,
		LoanCount: b.LoanCount,
	}
}
