package userparams

import (
	"lms-backend/internal/model"
	"lms-backend/internal/params/peopleparams"
	"lms-backend/pkg/error/externalerrors"
)

type UpdateParams struct {
	BaseUserParams
	ID           uint                      `json:"id"`
	PersonParams peopleparams.UpdateParams `json:"person_attributes"`
}

func (p *UpdateParams) ToModel() *model.User {
	usr := p.BaseUserParams.ToModel()
	usr.ID = p.ID
	usr.Person = p.PersonParams.ToModel()
	usr.PersonID = p.PersonParams.ID
	return usr
}

func (p *UpdateParams) Validate(userID int64) error {
	if p.ID == 0 {
		return externalerrors.BadRequest("User ID is required")
	}

	if p.ID != uint(userID) {
		return externalerrors.BadRequest("User ID does not match with the URL")
	}

	if err := p.BaseUserParams.Validate(); err != nil {
		return err
	}

	return p.PersonParams.Validate()
}
