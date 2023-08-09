package peopleparams

import (
	"technical-test/internal/model"

	"github.com/gofiber/fiber/v2"
)

type BaseParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (b *BaseParams) ToModel() *model.Person {
	return &model.Person{
		FirstName: b.FirstName,
		LastName:  b.LastName,
	}
}

func (b *BaseParams) Validate() error {
	if b.FirstName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "first_name is required")
	}

	if b.LastName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "last_name is required")
	}

	return nil
}
