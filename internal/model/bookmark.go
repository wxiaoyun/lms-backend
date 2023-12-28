package model

import (
	"fmt"
	"lms-backend/pkg/error/externalerrors"

	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model

	UserID uint  `gorm:"not null"`
	User   *User `gorm:"->"`
	BookID uint  `gorm:"not null"`
	Book   *Book `gorm:"->"`
}

const (
	BookmarkModelName = "bookmark"
	BookmarkTableName = "bookmarks"
)

const (
	MaximumBookmarkPerUser = 30
)

func (b *Bookmark) Create(db *gorm.DB) error {
	return db.Create(b).Error
}

func (b *Bookmark) Delete(db *gorm.DB) error {
	return db.Delete(b).Error
}

func (b *Bookmark) ensureUserExistsAndPresent(db *gorm.DB) error {
	if b.UserID == 0 {
		return externalerrors.BadRequest("user id is required")
	}

	var exists int64
	result := db.Model(&User{}).Where("id = ?", b.UserID).Count(&exists)
	if err := result.Error; err != nil {
		return err
	}

	if exists == 0 {
		return externalerrors.BadRequest("user does not exist")
	}

	return nil
}

func (b *Bookmark) ensureBookExistsAndPresent(db *gorm.DB) error {
	if b.BookID == 0 {
		return externalerrors.BadRequest("book id is required")
	}

	var exists int64
	result := db.Model(&Book{}).Where("id = ?", b.BookID).Count(&exists)
	if err := result.Error; err != nil {
		return err
	}

	if exists == 0 {
		return externalerrors.BadRequest("book does not exist")
	}

	return nil
}

func (b *Bookmark) ensureUserBookIsUnique(db *gorm.DB) error {
	var exists int64
	result := db.Model(&Bookmark{}).
		Where("book_id = ?", b.BookID).
		Where("user_id = ?", b.UserID).
		Count(&exists)
	if err := result.Error; err != nil {
		return err
	}

	if exists > 0 {
		return externalerrors.BadRequest("this bookmark already exists")
	}

	return nil
}

func (b *Bookmark) ensureMaxBookmarkPerUser(db *gorm.DB) error {
	var count int64
	result := db.Model(&Bookmark{}).
		Where("user_id = ?", b.UserID).
		Count(&count)
	if err := result.Error; err != nil {
		return err
	}

	if count >= MaximumBookmarkPerUser {
		return externalerrors.BadRequest(fmt.Sprintf("maximum bookmark per user is %d", MaximumBookmarkPerUser))
	}

	return nil
}

func (b *Bookmark) Validate(db *gorm.DB) error {
	if err := b.ensureUserExistsAndPresent(db); err != nil {
		return err
	}

	if err := b.ensureBookExistsAndPresent(db); err != nil {
		return err
	}

	if err := b.ensureUserBookIsUnique(db); err != nil {
		return err
	}

	return b.ensureMaxBookmarkPerUser(db)
}

func (b *Bookmark) BeforeCreate(tx *gorm.DB) error {
	return b.Validate(tx)
}
