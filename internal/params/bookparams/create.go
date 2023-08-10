package bookparams

import (
	"lms-backend/internal/model"
)

type CreateParams struct {
	BaseParams
}

func (p *CreateParams) Validate() error {
	return p.BaseParams.Validate()
}

func (p *CreateParams) ToModel() *model.Book {
	return p.BaseParams.ToModel()
}
