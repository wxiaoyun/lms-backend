package userview

import (
	"technical-test/internal/model"
	"technical-test/internal/view/personview"
)

type View struct {
	ID         uint            `json:"id,omitempty"`
	Email      string          `json:"email"`
	PersonView personview.View `json:"person_attributes"`
}

func ToView(user *model.User) *View {
	return &View{
		ID:         user.ID,
		Email:      user.Email,
		PersonView: *personview.ToView(user.Person),
	}
}
