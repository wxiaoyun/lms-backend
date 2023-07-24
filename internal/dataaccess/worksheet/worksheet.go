package worksheet

import (
	"technical-test/internal/model"

	"gorm.io/gorm"
)

func List(db *gorm.DB) ([]model.Worksheet, error) {
	var worksheets []model.Worksheet

	result := db.Model(&model.Worksheet{}).
		Find(&worksheets)
	if result.Error != nil {
		return nil, result.Error
	}

	return worksheets, nil
}
