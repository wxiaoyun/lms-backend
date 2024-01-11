package filestorage

import (
	"fmt"
	"lms-backend/pkg/error/externalerrors"
	"mime/multipart"
	"path/filepath"
	"strings"
)

const (
	megabyte    = 1_000_000
	maxFileSize = 10 * megabyte
)

var allowedContentTypes = map[string][]string{
	"image/jpeg":    {".jpeg", ".jpg", ".jfif", ".JPG"},
	"image/png":     {".png"},
	"image/gif":     {".gif"},
	"image/svg+xml": {".svg"},
	"image/webp":    {".webp"},
}

// isLessThanMaxFileSize checks if the uploaded file size is within the allowed limit.
func isLessThanMaxFileSize(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > maxFileSize {
		return externalerrors.BadRequest(
			fmt.Sprintf("file size %d exceeds maximum limit of %d bytes", fileHeader.Size, maxFileSize),
		)
	}

	return nil
}

// isAllowedContentType checks if the uploaded file's content type is allowed.
func isAllowedContentType(fileHeader *multipart.FileHeader) error {
	contentType := fileHeader.Header.Get("Content-Type")
	if _, ok := allowedContentTypes[contentType]; !ok {
		return externalerrors.BadRequest(
			fmt.Sprintf("%s is not an acceptable content type", contentType),
		)
	}

	return nil
}

// extensionMatchesContentType checks if the file extension matches the content type.
func extensionMatchesContentType(fileHeader *multipart.FileHeader) error {
	contentType := fileHeader.Header.Get("Content-Type")
	fileExtension := strings.ToLower(filepath.Ext(fileHeader.Filename))

	allowedExtensions, ok := allowedContentTypes[contentType]
	if !ok {
		return externalerrors.BadRequest(
			fmt.Sprintf("%s is not an acceptable content type", contentType),
		)
	}

	for _, ext := range allowedExtensions {
		if ext == fileExtension {
			return nil
		}
	}
	return externalerrors.BadRequest(
		fmt.Sprintf("file extension %s does not match content type %s", fileExtension, contentType),
	)
}

func ValidateFileUpload(fileHeader *multipart.FileHeader) error {
	if err := isLessThanMaxFileSize(fileHeader); err != nil {
		return err
	}

	if err := isAllowedContentType(fileHeader); err != nil {
		return err
	}

	return extensionMatchesContentType(fileHeader)
}
