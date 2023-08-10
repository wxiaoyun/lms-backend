package book

import (
	"lms-backend/internal/model"
	"lms-backend/internal/orm"

	"gorm.io/gorm"
)

func Read(db *gorm.DB, bookID int64) (*model.Book, error) {
	var book model.Book
	result := db.Model(&model.Book{}).
		Where("id = ?", bookID).
		First(&book)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.RecordNotFound(model.BookModelName)
		}
		return nil, err
	}

	return &book, nil
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

func Delete(db *gorm.DB, bookID int64) error {
	book, err := Read(db, bookID)
	if err != nil {
		return err
	}

	return book.Delete(db)
}
