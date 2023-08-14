package fine

import (
	"lms-backend/internal/model"
	"lms-backend/internal/orm"

	"gorm.io/gorm"
)

func Read(db *gorm.DB, fineID int64) (*model.Fine, error) {
	var fine model.Fine
	result := db.Model(&model.Fine{}).
		Where("id = ?", fineID).
		First(&fine)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.FineModelName)
		}
		return nil, err
	}

	return &fine, nil
}

func Create(db *gorm.DB, userID, loanID int64, amount float64) (*model.Fine, error) {
	fine := &model.Fine{
		UserID: uint(userID),
		LoanID: uint(loanID),
		Amount: amount,
		Status: model.FineStatusOutstanding,
	}

	if err := fine.Create(db); err != nil {
		return nil, err
	}

	return fine, nil
}

func Update(db *gorm.DB, fine *model.Fine) (*model.Fine, error) {
	if err := fine.Update(db); err != nil {
		return nil, err
	}

	return Read(db, int64(fine.ID))
}

func Delete(db *gorm.DB, fineID int64) (*model.Fine, error) {
	fn, err := Read(db, fineID)
	if err != nil {
		return nil, err
	}

	if err := fn.Delete(db); err != nil {
		return nil, err
	}

	return fn, nil
}

func Settle(db *gorm.DB, fineID int64) (*model.Fine, error) {
	fn, err := Read(db, fineID)
	if err != nil {
		return nil, err
	}

	fn.Status = model.FineStatusPaid

	if err := fn.Update(db); err != nil {
		return nil, err
	}

	return fn, nil
}
