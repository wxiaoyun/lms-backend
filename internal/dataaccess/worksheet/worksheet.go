package worksheet

import (
	"errors"
	"fmt"
	"technical-test/internal/model"
	viewmodel "technical-test/internal/viewmodel/worksheet"
	collection "technical-test/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.Preload("Questions")
}

func List(db *gorm.DB, cq *collection.Query) ([]model.Worksheet, error) {
	var worksheets []model.Worksheet

	result := db.Model(&model.Worksheet{}).
		Where("title ILIKE ? OR description ILIKE ?", "%"+cq.Search+"%", "%"+cq.Search+"%").
		Offset(cq.Offset).
		Limit(cq.Limit).
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
			return nil, fiber.NewError(fiber.StatusNotFound,
				fmt.Sprintf("worksheet %d is not found", id),
			)
		}
		return nil, err
	}

	return &worksheet, nil
}

func Count(db *gorm.DB) (int64, error) {
	var count int64

	result := db.Model(&model.Worksheet{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func CountFiltered(db *gorm.DB, cq *collection.Query) (int64, error) {
	var count int64

	result := db.Model(&model.Worksheet{}).
		Where("title ILIKE ? OR description ILIKE ?", "%"+cq.Search+"%", "%"+cq.Search+"%").
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func Summarize(db *gorm.DB) (*viewmodel.WorksheetSummaryViewModel, error) {
	var worksheets []model.Worksheet

	result := db.Model(&model.Worksheet{}).
		Scopes(preloadAssociations).
		Find(&worksheets)
	if err := result.Error; err != nil {
		return nil, err
	}

	var summary viewmodel.WorksheetSummaryViewModel
	for _, w := range worksheets {
		summary.TotalCost += w.GetTotalCost()
		summary.TotalPrice += w.GetTotalPrice()
		summary.TotalProfit += w.GetTotalProfit()
		if w.IsNegativeProfit() {
			summary.NegativeProfitCount++
		}
		if w.IsPositiveProfit() {
			summary.PositiveProfitCount++
		}
	}

	return &summary, nil
}
