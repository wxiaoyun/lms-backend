package userview

import (
	"auth-practice/internal/model"
)

type UserView struct {
	ID    uint   `json:"id,omitempty"`
	Email string `json:"email"`
}

func ToView(user *model.User) *UserView {
	return &UserView{
		ID:    user.ID,
		Email: user.Email,
	}
}
