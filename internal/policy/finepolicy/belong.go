package finepolicy

import (
	"lms-backend/internal/database"
	"lms-backend/internal/model"
	"lms-backend/internal/policy"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

type FineBelongsToUser struct {
	LoanID int64
	FineID int64
}

func AllowIfFineBelongsToUser(loanID, fineID int64) *FineBelongsToUser {
	return &FineBelongsToUser{
		LoanID: loanID,
		FineID: fineID,
	}
}

func (p *FineBelongsToUser) Validate(c *fiber.Ctx) (policy.Decision, error) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return policy.Deny, err
	}

	db := database.GetDB()

	var exists int64
	result := db.Model(&model.Fine{}).
		Joins("JOIN loans ON loans.id = fines.loan_id").
		Where("fines.id = ? AND fines.loan_id = ? AND loans.user_id = ?", p.FineID, p.LoanID, userID).
		Count(&exists)
	if result.Error != nil {
		return policy.Deny, result.Error
	}

	if exists == 0 {
		return policy.Deny, nil
	}

	return policy.Allow, nil
}
