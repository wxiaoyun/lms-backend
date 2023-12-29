package peopleparams

import (
	"lms-backend/internal/model"
	"lms-backend/pkg/error/externalerrors"

	"gorm.io/gorm"
)

type UpdateParams struct {
	BaseParams
	ID uint `json:"id"`
}

func (p *UpdateParams) ToModel() *model.Person {
	return &model.Person{
		Model: gorm.Model{
			ID: p.ID,
		},
		FullName:      p.FullName,
		PreferredName: p.PreferredName,
	}
}

func (p *UpdateParams) Validate() error {
	if err := p.BaseParams.Validate(); err != nil {
		return err
	}

	if p.ID == 0 {
		return externalerrors.BadRequest("Person ID is required")
	}

	return nil
}
