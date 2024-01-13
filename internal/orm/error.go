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

func IsDuplicatedKey(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey)
}

func IsForeignKeyViolated(err error) bool {
	return errors.Is(err, gorm.ErrForeignKeyViolated)
}

func ErrRecordNotFound(modelName string) error {
	return fiber.NewError(
		fiber.StatusBadRequest,
		fmt.Sprintf("%s not found", modelName),
	)
}
