package userparams

import "auth-practice/internal/model"

type BaseUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (b *BaseUserParams) ToModel() *model.User {
	return &model.User{
		Email:    b.Email,
		Password: b.Password,
	}
}
