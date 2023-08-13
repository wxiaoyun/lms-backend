package externalerrors

import (
	"github.com/gofiber/fiber/v2"
)

func BadRequest(message string) error {
	return fiber.NewError(fiber.StatusBadRequest, message)
}
