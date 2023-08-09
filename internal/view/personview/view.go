package personview

import (
	"technical-test/internal/model"
)

type View struct {
	ID        uint   `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func ToView(person *model.Person) *View {
	return &View{
		ID:        person.ID,
		FirstName: person.FirstName,
		LastName:  person.LastName,
	}
}
