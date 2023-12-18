package userparams

import (
	"lms-backend/internal/model"
	"lms-backend/pkg/error/externalerrors"
)

type SignInParams struct {
	BaseUserParams
}

func (p *SignInParams) ToModel() *model.User {
	return p.BaseUserParams.ToModel()
}

func (p *SignInParams) Validate() error {
	if err := p.BaseUserParams.Validate(); err != nil {
		return err
	}

	if p.Password == "" {
		return externalerrors.BadRequest("Password is required")
	}

	if p.Email == "" && p.Username == "" {
		return externalerrors.BadRequest("Email or Username is required")
	}

	return nil
}
