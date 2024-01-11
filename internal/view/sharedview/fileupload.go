package sharedview

import (
	"lms-backend/internal/model"
)

type FileUploadView struct {
	ID          uint   `json:"id,omitempty"`
	FileName    string `json:"file_name"`
	FilePath    string `json:"file_path"`
	ContentType string `json:"content_type"`
}

func ToFileUploadView(fileUpload *model.FileUpload) *FileUploadView {
	return &FileUploadView{
		ID:          fileUpload.ID,
		FileName:    fileUpload.FileName,
		FilePath:    fileUpload.FilePath,
		ContentType: fileUpload.ContentType,
	}
}
