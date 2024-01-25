package fileuploadhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/filestorage"
	"lms-backend/internal/model"
	"lms-backend/internal/session"
	"lms-backend/internal/view/sharedview"

	"github.com/gofiber/fiber/v2"
)

const (
	UploadField = "file"
)

func HandleUploadImage(c *fiber.Ctx) error {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	fileHeader, err := filestorage.ReadFileFromRequest(c, UploadField)
	if err != nil {
		return err
	}

	fileName, filePath, err := filestorage.SaveFileToDisk(c, fileHeader)
	if err != nil {
		return err
	}
	defer func() { // Delete file if panic occurs
		if r := recover(); r != nil || err != nil {
			//nolint
			filestorage.DeleteFileFromDisk(filePath)
			return
		}
	}()

	db := database.GetDB()
	username, err := user.GetUserName(db, userID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("User %s uploading thumbnail %s.", username, fileName),
	)
	defer func() { rollBackOrCommit(err) }()

	uploadModel := &model.FileUpload{
		FileName:    fileName,
		FilePath:    filePath,
		ContentType: fileHeader.Header.Get("Content-Type"),
	}

	err = uploadModel.Create(tx)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: sharedview.ToFileUploadView(uploadModel),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"\"%s\" uploaded successfully.", uploadModel.FileName,
			)),
		),
	})
}
