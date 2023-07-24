package userparams

import "technical-test/internal/model"

type CreateUserParams struct {
	BaseUserParams
}

func (c *CreateUserParams) ToModel() *model.User {
	return c.BaseUserParams.ToModel()
}

func (c *CreateUserParams) Validate() error {
	return c.BaseUserParams.Validate()
}
