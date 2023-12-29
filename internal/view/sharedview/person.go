package sharedview

import (
	"lms-backend/internal/model"
)

type PersonView struct {
	ID            uint   `json:"id,omitempty"`
	FullName      string `json:"full_name"`
	PreferredName string `json:"preferred_name"`
}

func ToPersonView(person *model.Person) *PersonView {
	return &PersonView{
		ID:            person.ID,
		FullName:      person.FullName,
		PreferredName: person.PreferredName,
	}
}
