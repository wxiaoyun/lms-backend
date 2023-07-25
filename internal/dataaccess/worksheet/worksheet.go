package worksheet

import (
	"technical-test/internal/model"

	"gorm.io/gorm"
)

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.Preload("Questions")
}

func List(db *gorm.DB) ([]model.Worksheet, error) {
	var worksheets []model.Worksheet

	result := db.Model(&model.Worksheet{}).
		Find(&worksheets)
	if result.Error != nil {
		return nil, result.Error
	}

	return worksheets, nil
}

func Read(db *gorm.DB, id int64) (*model.Worksheet, error) {
	var worksheet model.Worksheet

	result := db.Model(&model.Worksheet{}).
		Scopes(preloadAssociations).
		Where("id = ?", id).
		First(&worksheet)
	if result.Error != nil {
		return nil, result.Error
	}

	return &worksheet, nil
}
