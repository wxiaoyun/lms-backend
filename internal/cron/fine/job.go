package finejob

import (
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/model"
)

func DetectOverdueLoansAndCreateFine() {
	var err error

	tx, rollBackOrCommit := audit.Begin(nil, "CRON Job: Detecting overdue loans and creating fines")
	defer func() { rollBackOrCommit(err) }()

	var overdueLoansWithoutFines []model.Loan
	result := tx.Model(&model.Loan{}).
		Joins("LEFT JOIN fines ON loans.id = fines.loan_id").
		Where("fines.id IS NULL AND loans.due_date < NOW()").
		Where("loans.status = ?", model.LoanStatusBorrowed).
		Where("loans.return_date IS NULL").
		Find(&overdueLoansWithoutFines)
	if result.Error != nil {
		return
	}

	var fines = make([]model.Fine, 0, len(overdueLoansWithoutFines))
	for _, loan := range overdueLoansWithoutFines {
		fines = append(fines, model.Fine{
			UserID: loan.UserID,
			LoanID: loan.ID,
			Amount: model.OverdueFine,
			Status: model.FineStatusOutstanding,
		})
	}

	if len(fines) == 0 {
		return
	}

	if err = tx.Create(fines).Error; err != nil {
		return
	}
}
