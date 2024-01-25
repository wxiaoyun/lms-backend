package bookhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/filestorage"
	"lms-backend/internal/model"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"
	"lms-backend/internal/session"
	"lms-backend/internal/view/bookview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	UploadField = "file"
)

func HandleUpdateThumbnail(c *fiber.Ctx) error {
	err := policy.Authorize(c, updateBookAction, bookpolicy.UpdatePolicy())
	if err != nil {
		return err
	}

	param := c.Params("book_id")
	bookID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	fileHeader, err := filestorage.ReadFileFromRequest(c, UploadField)
	if err != nil {
		return err
	}

	fileName, filePath, err := filestorage.SaveFileToDisk(c, fileHeader, model.ThumbnailFolder)
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

	bookTitle, err := book.GetBookTitle(db, bookID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("User %s uploading thumbnail %s for \"%s\"", username, fileName, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	uploadModel := &model.FileUpload{
		FileName:    fileName,
		FilePath:    filePath,
		ContentType: fileHeader.Header.Get("Content-Type"),
	}

	bookModel, err := book.CreateOrUpdateThumbnail(tx, bookID, uploadModel)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: bookview.ToDetailedView(bookModel),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"\"%s\" uploaded successfully for \"%s\".", fileName, bookTitle,
			)),
		),
	})
}
