package fileupload

import (
	"lms-backend/internal/filestorage"
	"lms-backend/internal/model"
	"lms-backend/internal/orm"

	"gorm.io/gorm"
)

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.Preload("FileUploadReferences")
}

func Read(db *gorm.DB, fileUploadID int64) (*model.FileUpload, error) {
	var fileUpload model.FileUpload
	result := db.Model(&model.FileUpload{}).
		Scopes(preloadAssociations).
		Where("id = ?", fileUploadID).
		First(&fileUpload)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.FileUploadModelName)
		}
		return nil, err
	}

	return &fileUpload, nil
}

func Create(db *gorm.DB, fileUpload *model.FileUpload) (*model.FileUpload, error) {
	if err := fileUpload.Create(db); err != nil {
		return nil, err
	}

	return fileUpload, nil
}

func Delete(db *gorm.DB, fileUploadID int64) (*model.FileUpload, error) {
	fileUpload, err := Read(db, fileUploadID)
	if err != nil {
		return nil, err
	}

	for _, fileUploadReference := range fileUpload.FileUploadReferences {
		if err := fileUploadReference.Delete(db); err != nil {
			return nil, err
		}
	}

	if err := fileUpload.Delete(db); err != nil {
		return nil, err
	}

	if err := filestorage.DeleteFileFromDisk(fileUpload.FilePath); err != nil {
		return nil, err
	}

	return fileUpload, nil
}
