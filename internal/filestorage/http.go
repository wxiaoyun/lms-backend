package filestorage

import (
	"fmt"
	"lms-backend/pkg/error/externalerrors"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

func ReadFileFromRequest(c *fiber.Ctx, field string) (*multipart.FileHeader, error) {
	fileHeader, err := c.FormFile(field)
	if err != nil {
		return nil, externalerrors.BadRequest(
			fmt.Sprintf("could not read file from field: %s", field),
		)
	}

	return fileHeader, nil
}
