package book

import (
	"lms-backend/internal/dataaccess/loan"
	"lms-backend/internal/dataaccess/reservation"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/model"
	"lms-backend/internal/orm"
	"lms-backend/internal/viewmodel"
	"lms-backend/pkg/error/externalerrors"
	"time"

	"gorm.io/gorm"
)

func preloadCopies(db *gorm.DB) *gorm.DB {
	return db.Preload("BookCopies")
}

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.Scopes(preloadCopies).
		Preload("Bookmarks").
		Preload("Thumbnail").
		Preload("Thumbnail.FileUpload")
}

func Read(db *gorm.DB, bookID int64) (*model.Book, error) {
	var book model.Book
	result := db.Model(&model.Book{}).
		Where("id = ?", bookID).
		First(&book)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.BookModelName)
		}
		return nil, err
	}

	return &book, nil
}

func ReadWithCopies(db *gorm.DB, bookID int64) (*model.Book, error) {
	var book model.Book
	result := db.Model(&model.Book{}).
		Scopes(preloadCopies).
		Where("id = ?", bookID).
		First(&book)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.BookModelName)
		}
		return nil, err
	}

	return &book, nil
}

func ReadDetailed(db *gorm.DB, bookID int64) (*model.Book, error) {
	var book model.Book
	result := db.Model(&model.Book{}).
		Scopes(preloadAssociations).
		Where("id = ?", bookID).
		First(&book)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.BookModelName)
		}
		return nil, err
	}

	return &book, nil
}

func GetBookTitle(db *gorm.DB, bookID int64) (string, error) {
	var title string

	result := db.Model(&model.Book{}).
		Select("title").
		Where("id = ?", bookID).
		First(&title)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return "", orm.ErrRecordNotFound(model.BookModelName)
		}
		return "", err
	}

	return title, nil
}

func Create(db *gorm.DB, book *model.Book) (*model.Book, error) {
	if err := book.Create(db); err != nil {
		return nil, err
	}

	return Read(db, int64(book.ID))
}

func CreateWithCopy(db *gorm.DB, book *model.Book) (*model.Book, error) {
	book.BookCopies = []model.BookCopy{{}} // initialize with one copy
	return Create(db, book)
}

func Update(db *gorm.DB, book *model.Book) (*model.Book, error) {
	if err := book.Update(db); err != nil {
		return nil, err
	}

	return book, nil
}

func Delete(db *gorm.DB, bookID int64) (*model.Book, error) {
	book, err := ReadDetailed(db, bookID)
	if err != nil {
		return nil, err
	}

	for _, copy := range book.BookCopies {
		if err := copy.Delete(db); err != nil {
			return nil, err
		}
	}

	for _, bookmark := range book.Bookmarks {
		if err := bookmark.Delete(db); err != nil {
			return nil, err
		}
	}

	if err := book.Delete(db); err != nil {
		return nil, err
	}

	return book, nil
}

func Count(db *gorm.DB) (int64, error) {
	var count int64

	result := orm.CloneSession(db).
		Model(&model.Book{}).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func List(db *gorm.DB) ([]model.Book, error) {
	var books []model.Book

	result := db.Model(&model.Book{}).
		Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func ListWithCopies(db *gorm.DB) ([]model.Book, error) {
	books, err := List(db)
	if err != nil {
		return nil, err
	}

	db = orm.NewSession(db)

	for i, b := range books {
		var copies []model.BookCopy

		result := db.Model(&model.BookCopy{}).
			Where("book_id = ?", b.ID).
			Find(&copies)
		if result.Error != nil {
			return nil, result.Error
		}

		books[i].BookCopies = copies
	}

	return books, nil
}

func AutoComplete(db *gorm.DB, value string) ([]model.Book, error) {
	if len(value) == 0 {
		return []model.Book{}, nil
	}

	var books []model.Book

	result := db.Model(&model.Book{}).
		Where("title ILIKE ?", "%%"+value+"%%").
		Or("author ILIKE ?", "%%"+value+"%%").
		Or("publisher ILIKE ?", "%%"+value+"%%").
		Or("isbn ILIKE ?", "%%"+value+"%%").
		Limit(5).
		Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func ListPopularBooks(db *gorm.DB) ([]viewmodel.BookLoanCount, error) {
	// Calculate the date 3 months ago from now
	threeMonthsAgo := time.Now().AddDate(0, -3, 0)

	var bookLoanCounts []viewmodel.BookLoanCount

	result := db.Model(&model.Book{}).
		Select("books.id, books.title, COUNT(loans.id) AS loan_count").
		Joins("INNER JOIN book_copies ON books.id = book_copies.book_id").
		Joins("INNER JOIN loans ON book_copies.id = loans.book_copy_id AND loans.created_at >= ?", threeMonthsAgo).
		Group("books.id, books.title").
		Order("loan_count DESC").
		Limit(10).
		Find(&bookLoanCounts)

	if result.Error != nil {
		return nil, result.Error
	}

	return bookLoanCounts, nil
}

// Ensure the same user cannot loan more than 1 copy of the same book
func CountNumberOfCopiesLoanedByUser(db *gorm.DB, userID, bookID int64) (int64, error) {
	var count int64

	result := db.Model(&model.BookCopy{}).
		Joins("INNER JOIN loans ON book_copies.id = loans.book_copy_id").
		Where("loans.status = ?", model.LoanStatusBorrowed).
		Where("loans.user_id = ?", userID).
		Where("book_copies.book_id = ?", bookID).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func CountNumberOfCopiesReservedByUser(db *gorm.DB, userID, bookID int64) (int64, error) {
	var count int64

	result := db.Model(&model.BookCopy{}).
		Joins("INNER JOIN reservations ON book_copies.id = reservations.book_copy_id").
		Where("reservations.status = ?", model.ReservationStatusPending).
		Where("reservations.user_id = ?", userID).
		Where("book_copies.book_id = ?", bookID).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func Loan(db *gorm.DB, userID, bookID int64) (*model.Loan, error) {
	hasExceededMaxLoan, err := user.HasExceededMaxLoan(db, userID)
	if err != nil {
		return nil, err
	}
	if hasExceededMaxLoan {
		return nil, externalerrors.BadRequest("You have reached the maximum number of loans")
	}

	loanCount, err := CountNumberOfCopiesLoanedByUser(db, userID, bookID)
	if err != nil {
		return nil, err
	}
	if loanCount > 0 {
		return nil, externalerrors.BadRequest("You have already loaned a copy of this book")
	}

	resCount, err := CountNumberOfCopiesReservedByUser(db, userID, bookID)
	if err != nil {
		return nil, err
	}
	if resCount > 0 {
		return nil, externalerrors.BadRequest("You have already reserved a copy of this book")
	}

	book, err := ReadWithCopies(db, bookID)
	if err != nil {
		return nil, err
	}

	for _, copy := range book.BookCopies {
		if copy.Status == model.BookStatusOnLoan {
			continue
		}

		if copy.Status == model.BookStatusOnReserve {
			continue
		}

		ln, err := loan.Loan(db, userID, int64(copy.ID))
		if err != nil {
			return nil, err
		}

		// Update book status
		copy.Status = model.BookStatusOnLoan
		if err := copy.Update(db); err != nil {
			return nil, err
		}

		return ln, nil
	}

	return nil, externalerrors.BadRequest("No copies are available for loan")
}

func Reserve(db *gorm.DB, userID, bookID int64) (*model.Reservation, error) {
	hasExceededMaxReservation, err := user.HasExceededMaxReservation(db, userID)
	if err != nil {
		return nil, err
	}
	if hasExceededMaxReservation {
		return nil, externalerrors.BadRequest("You have reached the maximum number of reservations")
	}

	loanCount, err := CountNumberOfCopiesLoanedByUser(db, userID, bookID)
	if err != nil {
		return nil, err
	}
	if loanCount > 0 {
		return nil, externalerrors.BadRequest("You have already loaned a copy of this book")
	}

	resCount, err := CountNumberOfCopiesReservedByUser(db, userID, bookID)
	if err != nil {
		return nil, err
	}
	if resCount > 0 {
		return nil, externalerrors.BadRequest("You have already reserved a copy of this book")
	}

	book, err := ReadWithCopies(db, bookID)
	if err != nil {
		return nil, err
	}

	for _, copy := range book.BookCopies {
		if copy.Status == model.BookStatusOnLoan {
			continue
		}

		if copy.Status == model.BookStatusOnReserve {
			continue
		}

		res, err := reservation.ReserveBook(db, userID, int64(copy.ID))
		if err != nil {
			return nil, err
		}

		// Update book status
		copy.Status = model.BookStatusOnReserve
		if err := copy.Update(db); err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, externalerrors.BadRequest("No copies are available for reservation")
}
