package book

import (
	"lms-backend/internal/model"
	"lms-backend/internal/orm"

	"gorm.io/gorm"
)

func preloadCopies(db *gorm.DB) *gorm.DB {
	return db.Preload("BookCopies")
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
	book, err := Read(db, bookID)
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
