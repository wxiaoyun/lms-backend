package worksheet

import (
	"errors"
	"technical-test/internal/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.Preload("Questions")
}

func List(db *gorm.DB) ([]model.Worksheet, error) {
	var worksheets []model.Worksheet

	result := db.Model(&model.Worksheet{}).
		Scopes(preloadAssociations).
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
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		return nil, err
	}

	return &worksheet, nil
}
