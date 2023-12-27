package bookmark

import (
	"lms-backend/internal/model"
	"lms-backend/internal/orm"

	"gorm.io/gorm"
)

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Books")
}

func Read(db *gorm.DB, id int64) (*model.Bookmark, error) {
	var b model.Bookmark
	result := db.Model(&model.Bookmark{}).
		Where("id = ?", id).
		First(&b)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.BookmarkModelName)
		}
		return nil, err
	}

	return &b, nil
}

func ReadDetailed(db *gorm.DB, id int64) (*model.Bookmark, error) {
	var b model.Bookmark
	result := db.Model(&model.Bookmark{}).
		Scopes(preloadAssociations).
		Where("id = ?", id).
		First(&b)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.BookmarkModelName)
		}
		return nil, err
	}

	return &b, nil
}

func Create(db *gorm.DB, b *model.Bookmark) (*model.Bookmark, error) {
	if err := b.Create(db); err != nil {
		return nil, err
	}

	return ReadDetailed(db, int64(b.ID))
}

func Delete(db *gorm.DB, id int64) (*model.Bookmark, error) {
	b, err := ReadDetailed(db, id)
	if err != nil {
		return nil, err
	}

	if err := b.Delete(db); err != nil {
		return nil, err
	}

	return b, nil
}

func Count(db *gorm.DB) (int64, error) {
	var count int64

	result := orm.CloneSession(db).
		Model(&model.Bookmark{}).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func List(db *gorm.DB) ([]*model.Bookmark, error) {
	var bookmarks []*model.Bookmark
	result := db.Model(&model.Bookmark{}).
		Find(&bookmarks)
	if err := result.Error; err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func ListDetailed(db *gorm.DB) ([]*model.Bookmark, error) {
	var bookmarks []*model.Bookmark
	result := db.Model(&model.Bookmark{}).
		Scopes(preloadAssociations).
		Find(&bookmarks)
	if err := result.Error; err != nil {
		return nil, err
	}

	return bookmarks, nil
}
