package userview

import (
	"lms-backend/internal/model"
)

type CurrentUserView struct {
	IsLoggedIn bool `json:"is_logged_in"`
	LoginView
}

func ToGuestView() *CurrentUserView {
	return &CurrentUserView{
		IsLoggedIn: false,
		LoginView: LoginView{
			Abilities: []string{},
		},
	}
}

func ToCurrentUserView(
	user *model.User,
	abilities []model.Ability,
) *CurrentUserView {
	return &CurrentUserView{
		IsLoggedIn: true,
		LoginView:  *ToLoginView(user, abilities),
	}
}
