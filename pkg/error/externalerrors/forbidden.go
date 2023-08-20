package externalerrors

import (
	"github.com/gofiber/fiber/v2"
)

func Forbidden(message string) error {
	return fiber.NewError(fiber.StatusForbidden, message)
}
