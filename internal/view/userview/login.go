package userview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/personview"
	"lms-backend/internal/view/sharedview"
	"lms-backend/util/sliceutil"
)

type LoginView struct {
	User       *sharedview.UserView `json:"user"`
	PersonView *personview.View     `json:"person_attributes"`
	Abilities  []string             `json:"abilities"`
}

func ToLoginView(
	user *model.User,
	abilities []model.Ability,
) *LoginView {
	return &LoginView{
		User:       sharedview.ToUserView(user),
		PersonView: personview.ToView(user.Person),
		Abilities:  sliceutil.Map(abilities, func(a model.Ability) string { return a.Name }),
	}
}
