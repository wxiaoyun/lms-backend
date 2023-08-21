package userview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/personview"
	"lms-backend/util/sliceutil"
)

type View struct {
	ID         uint            `json:"id,omitempty"`
	Username   string          `json:"username"`
	Email      string          `json:"email,omitempty"`
	PersonView personview.View `json:"person_attributes"`
	Abilities  []string        `json:"abilities,omitempty"`
}

func ToView(user *model.User, abilities []model.Ability) *View {
	return &View{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		PersonView: *personview.ToView(user.Person),
		Abilities:  sliceutil.Map(abilities, func(a model.Ability) string { return a.Name }),
	}
}
