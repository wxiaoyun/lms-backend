package book

import (
	"lms-backend/internal/dataaccess/fileupload"
	"lms-backend/internal/model"

	"gorm.io/gorm"
)

func CreateOrUpdateThumbnail(db *gorm.DB, bookID int64, thumbnail *model.FileUpload) (*model.Book, error) {
	// preload thumbnail and its file upload
	book, err := ReadDetailed(db, bookID)
	if err != nil {
		return nil, err
	}

	// delete old thumbnail if exists
	if book.Thumbnail != nil {
		_, err := fileupload.Delete(db, int64(book.Thumbnail.ID))
		if err != nil {
			return nil, err
		}
	}

	// create new thumbnail
	thumbnailRef := &model.FileUploadReference{
		FileUpload:     thumbnail,
		AttachableID:   book.ID,
		AttachableType: model.BookThumbnailFileUploadReferenceAttachableType,
	}

	if err := thumbnailRef.Create(db); err != nil {
		return nil, err
	}

	book.Thumbnail = thumbnailRef

	return book, nil
}
