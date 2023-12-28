package sharedview

import (
	"lms-backend/internal/model"
)

type UserView struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
}

func ToUserView(user *model.User) *UserView {
	return &UserView{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
