package filestorage

import (
	"lms-backend/internal/config"
	"lms-backend/pkg/storage"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/pkg/errors"
)

var Storage = storage.Storage{
	BaseDirectoryElems: []string{config.RuntimeWorkingDirectory, "file_storage"},
}

// SaveFileToDisk saves the file from the request to the disk.
// It returns the filename of the saved file and the path to the file.
//
//nolint:revive
func SaveFileToDisk(c *fiber.Ctx, fileHeader *multipart.FileHeader, subdirectory ...string) (string, string, error) {
	err := ValidateFileUpload(fileHeader)
	if err != nil {
		return "", "", err
	}

	fileUUID := utils.UUIDv4()
	filename := fileUUID + filepath.Ext(fileHeader.Filename)

	filePath, err := Storage.ConstructFilePath(append(subdirectory, filename)...)
	if err != nil {
		return "", "", err
	}

	// Create directory
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", "", errors.Wrap(err, "creating directory failed")
	}

	if err := c.SaveFile(fileHeader, filePath); err != nil {
		return "", "", err
	}

	return filename, filePath, nil
}

// filePath should not be a user input.
func DeleteFileFromDisk(filePath string) error {
	if err := Storage.ValidateFilePath(filePath); err != nil {
		return err
	}

	return os.Remove(filePath)
}

func FileExists(filePath string) bool {
	dir := filepath.Dir("file_storage/" + filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return false
	}

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
