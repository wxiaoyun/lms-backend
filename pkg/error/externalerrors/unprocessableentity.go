package externalerrors

import (
	"github.com/gofiber/fiber/v2"
)

func UnprocessableEntity(message string) error {
	return fiber.NewError(fiber.StatusUnprocessableEntity, message)
}
