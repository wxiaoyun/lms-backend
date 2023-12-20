package internalerror

import (
	"github.com/gofiber/fiber/v2"
)

func InternalServerError(message string) error {
	return fiber.NewError(fiber.StatusInternalServerError, message)
}
