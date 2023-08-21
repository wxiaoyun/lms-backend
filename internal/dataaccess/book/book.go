package book

import (
	"lms-backend/internal/dataaccess/loan"
	"lms-backend/internal/dataaccess/reservation"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/model"
	"lms-backend/internal/orm"
	collection "lms-backend/pkg/collectionquery"
	"lms-backend/pkg/error/externalerrors"

	"gorm.io/gorm"
)

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Loans").
		Preload("Loans.LoanHistories").
		Preload("Loans.Fines").
		Preload("Reservations")
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

func Update(db *gorm.DB, book *model.Book) (*model.Book, error) {
	if err := book.Update(db); err != nil {
		return nil, err
	}

	return Read(db, int64(book.ID))
}

func Delete(db *gorm.DB, bookID int64) (*model.Book, error) {
	book, err := ReadDetailed(db, bookID)
	if err != nil {
		return nil, err
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

func CountFiltered(db *gorm.DB, cq *collection.Query) (int64, error) {
	var count int64

	result := db.Model(&model.Book{}).
		Where("title ILIKE ?", "%"+cq.Search+"%").
		Or("author ILIKE ?", "%"+cq.Search+"%").
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

func IsOnLoan(db *gorm.DB, bookID int64) (bool, error) {
	var count int64

	result := db.Model(&model.Loan{}).
		Where("book_id = ?", bookID).
		Where("status = ?", model.LoanStatusBorrowed).
		Where("return_date IS NULL").
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func IsOnReserve(db *gorm.DB, bookID int64) (bool, error) {
	var count int64

	result := db.Model(&model.Reservation{}).
		Where("book_id = ?", bookID).
		Where("status = ?", model.ReservationStatusPending).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func LoanBook(db *gorm.DB, userID, bookID int64) (*model.Loan, error) {
	isOnLoan, err := IsOnLoan(db, bookID)
	if err != nil {
		return nil, err
	}
	if isOnLoan {
		return nil, externalerrors.BadRequest("Book is already on loan")
	}

	isOnReserve, err := IsOnReserve(db, bookID)
	if err != nil {
		return nil, err
	}
	if isOnReserve {
		return nil, externalerrors.BadRequest("Book is already on reserve")
	}

	hasExceededMaxLoan, err := user.HasExceededMaxLoan(db, userID)
	if err != nil {
		return nil, err
	}
	if hasExceededMaxLoan {
		return nil, externalerrors.BadRequest("You have reached the maximum number of loans")
	}

	// Create loan
	ln, err := loan.LoanBook(db, userID, bookID)
	if err != nil {
		return nil, err
	}

	return ln, nil
}

func ReturnBook(db *gorm.DB, loanID int64) (*model.Loan, error) {
	returnedLn, err := loan.ReturnBook(db, loanID)
	if err != nil {
		return nil, err
	}

	return returnedLn, nil
}

func RenewLoan(db *gorm.DB, loanID int64) (*model.Loan, error) {
	ln, err := loan.Read(db, loanID)
	if err != nil {
		return nil, err
	}

	reservations, err := reservation.ReadOutstandingReservationsByBookID(db, int64(ln.BookID))
	if err != nil {
		return nil, err
	}

	if len(reservations) > 0 {
		// check if the book is reserved by another user
		return nil, externalerrors.BadRequest("You cannot renew the loan because the book is reserved")
	}

	renewedLn, err := loan.RenewLoan(db, loanID)
	if err != nil {
		return nil, err
	}

	return renewedLn, nil
}

func ReserveBook(db *gorm.DB, userID, bookID int64) (*model.Reservation, error) {
	isOnReserve, err := IsOnReserve(db, bookID)
	if err != nil {
		return nil, err
	}
	if isOnReserve {
		return nil, externalerrors.BadRequest("Book is already on reserve")
	}

	hasExceededMaxReservation, err := user.HasExceededMaxReservation(db, userID)
	if err != nil {
		return nil, err
	}
	if hasExceededMaxReservation {
		return nil, externalerrors.BadRequest("You have reached the maximum number of reservations")
	}

	// Create reservation
	res, err := reservation.ReserveBook(db, userID, bookID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CheckOutReservation(db *gorm.DB, userID, bookID, resID int64) (*model.Reservation, error) {
	isOnLoan, err := IsOnLoan(db, bookID)
	if err != nil {
		return nil, err
	}
	if isOnLoan {
		return nil, externalerrors.BadRequest("Book is currently on loan")
	}

	// Fulfill the reservation
	res, err := reservation.FullfilReservation(db, resID)
	if err != nil {
		return nil, err
	}

	// Loan the book to the user
	_, err = LoanBook(db, userID, bookID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
