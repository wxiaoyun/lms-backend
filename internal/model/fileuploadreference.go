package model

import (
	"fmt"
	"lms-backend/internal/config"
	"lms-backend/pkg/error/externalerrors"

	"gorm.io/gorm"
)

type FileUploadReference struct {
	gorm.Model

	FileUploadID   uint        `gorm:"not null"`
	FileUpload     *FileUpload `gorm:"->;<-:create"`
	AttachableID   uint        `gorm:"not null"`
	AttachableType string      `gorm:"not null"`
}

const (
	BookThumbnailFileUploadReferenceAttachableType = "book_thumbnail"
)

const (
	ImageDownloadURL = "%s/file/image/%s"
)

func (f *FileUploadReference) Create(db *gorm.DB) error {
	return db.Create(f).Error
}

func (f *FileUploadReference) Delete(db *gorm.DB) error {
	return db.Delete(f).Error
}

func (f *FileUploadReference) ensureNoDuplicate(db *gorm.DB) error {
	var count int64
	err := db.Model(f).
		Where("file_upload_id = ? AND attachable_id = ? AND attachable_type = ?", f.FileUploadID, f.AttachableID, f.AttachableType).
		Count(&count).
		Error
	if err != nil {
		return err
	}

	if count > 0 {
		return externalerrors.BadRequest("file upload reference already exists")
	}

	return nil
}

func (f *FileUploadReference) ensureFileUploadExistsOrIsNew(db *gorm.DB) error {
	if f.FileUploadID == 0 {
		return nil
	}

	var count int64
	err := db.Model(f.FileUpload).
		Where("id = ?", f.FileUploadID).
		Count(&count).
		Error
	if err != nil {
		return err
	}

	if count == 0 {
		return externalerrors.BadRequest("file upload does not exist")
	}

	return nil
}

func (f *FileUploadReference) Validate(db *gorm.DB) error {
	if err := f.ensureNoDuplicate(db); err != nil {
		return err
	}

	return f.ensureFileUploadExistsOrIsNew(db)
}

func (f *FileUploadReference) BeforeCreate(db *gorm.DB) error {
	return f.Validate(db)
}

// FileUpload needs to be preloaded
func (f *FileUploadReference) GetImageDownloadURL() string {
	return fmt.Sprintf(ImageDownloadURL, config.BackendURL, f.FileUpload.FileName)
}
