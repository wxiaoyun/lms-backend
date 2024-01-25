package model

import (
	"lms-backend/pkg/error/externalerrors"

	"gorm.io/gorm"
)

type FileUpload struct {
	gorm.Model

	FileName             string                `gorm:"not null"`
	FilePath             string                `gorm:"not null"`
	ContentType          string                `gorm:"not null"`
	FileUploadReferences []FileUploadReference `gorm:"->"`
}

const (
	FileUploadModelName = "file_upload"
)

const (
	ThumbnailFolder = BookThumbnailFileUploadReferenceAttachableType
	ImageFolder     = "image"
	QRCodeFolder    = "qrcode"
)

func (f *FileUpload) Create(db *gorm.DB) error {
	return db.Create(f).Error
}

func (f *FileUpload) Delete(db *gorm.DB) error {
	return db.Delete(f).Error
}

func (f *FileUpload) Validate(_ *gorm.DB) error {
	if f.FileName == "" {
		return externalerrors.BadRequest("file name is required")
	}

	if f.FilePath == "" {
		return externalerrors.BadRequest("file path is required")
	}

	if f.ContentType == "" {
		return externalerrors.BadRequest("content type is required")
	}

	return nil
}

func (f *FileUpload) BeforeCreate(_ *gorm.DB) error {
	return f.Validate(nil)
}
