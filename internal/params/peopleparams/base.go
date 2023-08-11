package peopleparams

import (
	"lms-backend/internal/model"

	"github.com/gofiber/fiber/v2"
)

type BaseParams struct {
	FullName           string `json:"full_name"`
	PreferredName      string `json:"preferred_name"`
	LanguagePreference string `json:"language_preference"`
}

func (b *BaseParams) ToModel() *model.Person {
	return &model.Person{
		FullName:           b.FullName,
		PreferredName:      b.PreferredName,
		LanguagePreference: b.LanguagePreference,
	}
}

func (b *BaseParams) Validate() error {
	if b.FullName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "full_name is required")
	}

	return nil
}
