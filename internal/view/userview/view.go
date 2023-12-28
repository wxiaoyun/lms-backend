package userview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/personview"
	"lms-backend/internal/view/sharedview"
	"lms-backend/util/sliceutil"
)

type View struct {
	sharedview.UserView
	PersonView *personview.View `json:"person_attributes"`
	Abilities  []string         `json:"abilities,omitempty"`
}

func ToView(user *model.User, abilities ...model.Ability) *View {
	return &View{
		UserView:   *sharedview.ToUserView(user),
		PersonView: personview.ToView(user.Person),
		Abilities:  sliceutil.Map(abilities, func(a model.Ability) string { return a.Name }),
	}
}
