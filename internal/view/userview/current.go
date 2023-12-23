package userview

import (
	"lms-backend/internal/model"
)

type CurrentUserView struct {
	IsLoggedIn bool  `json:"is_logged_in"`
	User       *View `json:"user"`
}

func ToCurrentUserView(user *model.User, abilities ...model.Ability) *CurrentUserView {
	if user == nil {
		return &CurrentUserView{
			IsLoggedIn: false,
		}
	}

	return &CurrentUserView{
		IsLoggedIn: true,
		User:       ToView(user, abilities...),
	}
}
