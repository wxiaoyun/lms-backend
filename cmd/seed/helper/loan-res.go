package shelper

import (
	"lms-backend/internal/dataaccess/bookcopy"
	"lms-backend/internal/model"
	"lms-backend/util/random"
	"math/rand"

	"gorm.io/gorm"
)

func SeedLoanAndReservations(db *gorm.DB) error {
	var bookNum int64
	result := db.Model(&model.BookCopy{}).Count(&bookNum)
	if result.Error != nil {
		return result.Error
	}

	var userNum int64
	result = db.Model(&model.User{}).Count(&userNum)
	if result.Error != nil {
		return result.Error
	}

	ids := make([]int64, bookNum)
	for i := 1; i <= int(bookNum); i++ {
		ids[i-1] = int64(i)
	}

	// Shuffle IDs
	rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })

	for userID := 1; userID <= int(userNum); userID++ {
		for j := 1; j <= random.RandInt(0, model.MaximumLoans+1); j++ {
			if len(ids) == 0 {
				return nil
			}

			copyID := ids[0]
			ids = ids[1:]

			_, err := bookcopy.LoanCopy(db, int64(userID), copyID)
			if err != nil {
				return err
			}
		}

		for j := 1; j <= random.RandInt(0, model.MaximumReservations+1); j++ {
			if len(ids) == 0 {
				return nil
			}

			copyID := ids[0]
			ids = ids[1:]

			_, err := bookcopy.ReserveCopy(db, int64(userID), copyID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
