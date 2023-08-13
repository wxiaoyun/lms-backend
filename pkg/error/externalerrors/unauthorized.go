package externalerrors

import (
	"github.com/gofiber/fiber/v2"
)

func Unauthorized(message string) error {
	return fiber.NewError(fiber.StatusUnauthorized, message)
}
