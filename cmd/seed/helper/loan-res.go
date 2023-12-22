package shelper

import (
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/model"
	"math/rand"

	"gorm.io/gorm"
)

func SeedLoanAndReservations(db *gorm.DB, userNum, bookNum int64) error {
	ids := make([]int64, bookNum)
	for i := 1; i <= int(bookNum); i++ {
		ids[i-1] = int64(i)
	}

	// Shuffle IDs
	rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })

	for userID := 1; userID <= int(userNum); userID++ {
		for j := 1; j <= model.MaximumLoans; j++ {
			if len(ids) == 0 {
				return nil
			}

			bookID := ids[0]
			ids = ids[1:]

			_, err := book.LoanBook(db, int64(userID), bookID)
			if err != nil {
				return err
			}
		}

		for j := 1; j <= model.MaximumReservations; j++ {
			if len(ids) == 0 {
				return nil
			}

			bookID := ids[0]
			ids = ids[1:]

			_, err := book.ReserveBook(db, int64(userID), bookID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
