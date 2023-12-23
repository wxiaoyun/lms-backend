package shelper

import (
	"lms-backend/internal/model"
	"lms-backend/util/random"
	"time"

	"gorm.io/gorm"
)

func SeedBooks(db *gorm.DB, num int64) error {
	var count int64

	result := db.Model(&model.Book{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count >= num {
		return nil
	}

	books := make([]model.Book, num)
	for i := 1; i <= int(num); i++ {
		books[i-1] = model.Book{
			Title:           random.RandWords(random.RandInt(4, 11)),
			Author:          random.RandWords(random.RandInt(2, 5)),
			ISBN:            GenerateISBN13(),
			Publisher:       random.RandWords(random.RandInt(4, 7)),
			PublicationDate: random.RandomDate(time.Now().AddDate(-10, 0, 0), time.Now()),
			Genre:           random.RandWords(random.RandInt(1, 3)),
			Language:        random.RandWords(random.RandInt(1, 3)),
		}
	}

	return db.Create(&books).Error
}
