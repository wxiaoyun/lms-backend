package userparams

import "auth-practice/internal/model"

type CreateUserParams struct {
	BaseUserParams
}

func (c *CreateUserParams) ToModel() *model.User {
	return c.BaseUserParams.ToModel()
}
