package bookparams

import (
	"lms-backend/internal/model"

	"github.com/gofiber/fiber/v2"
)

type UpdateParams struct {
	ID uint `json:"id"`
	BaseParams
}

func (p *UpdateParams) Validate() error {
	if p.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	return p.BaseParams.Validate()
}

func (p *UpdateParams) ToModel() *model.Book {
	book := p.BaseParams.ToModel()
	book.ID = p.ID
	return book
}
