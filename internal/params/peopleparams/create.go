package peopleparams

import (
	"lms-backend/internal/model"

	"github.com/gofiber/fiber/v2"
)

type CreateParams struct {
	FullName      string `json:"full_name"`
	PreferredName string `json:"preferred_name"`
}

func (b *CreateParams) ToModel() *model.Person {
	return &model.Person{
		FullName:           b.FullName,
		PreferredName:      b.PreferredName,
		LanguagePreference: "EN",
	}
}

func (b *CreateParams) Validate() error {
	if b.FullName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "full_name is required")
	}

	return nil
}
