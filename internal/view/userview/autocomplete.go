package userview

import (
	"lms-backend/internal/model"
)

type SimpleView struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func ToSimpleView(user *model.User) *SimpleView {
	return &SimpleView{
		ID:       user.ID,
		Username: user.Username,
	}
}
