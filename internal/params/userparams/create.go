package userparams

import (
	"lms-backend/internal/model"
	"lms-backend/internal/params/peopleparams"
)

type CreateParams struct {
	BaseUserParams
	PersonParams peopleparams.CreateParams `json:"person_attributes"`
}

func (p *CreateParams) ToModel() *model.User {
	usr := p.BaseUserParams.ToModel()
	usr.Person = p.PersonParams.ToModel()
	return usr
}

func (p *CreateParams) Validate() error {
	if err := p.BaseUserParams.Validate(); err != nil {
		return err
	}

	return p.PersonParams.Validate()
}
