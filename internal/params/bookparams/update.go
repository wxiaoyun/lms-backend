package bookparams

import (
	"lms-backend/internal/model"
	"lms-backend/pkg/error/externalerrors"
)

type UpdateParams struct {
	ID uint `json:"id"`
	BaseParams
}

func (p *UpdateParams) Validate(bookID int64) error {
	if p.ID == 0 {
		return externalerrors.BadRequest("id is required")
	}

	if p.ID != uint(bookID) {
		return externalerrors.BadRequest("book ID is inconsistent with url")
	}

	return p.BaseParams.Validate()
}

func (p *UpdateParams) ToModel() *model.Book {
	book := p.BaseParams.ToModel()
	book.ID = p.ID
	return book
}
