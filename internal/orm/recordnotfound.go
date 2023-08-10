package orm

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func RecordNotFound(modelName string) error {
	return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("%s not found", modelName))
}
