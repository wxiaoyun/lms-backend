package personview

import (
	"lms-backend/internal/model"
)

type View struct {
	ID                 uint   `json:"id,omitempty"`
	FullName           string `json:"full_name"`
	PreferredName      string `json:"preferred_name"`
	LanguagePreference string `json:"language_preference"`
}

func ToView(person *model.Person) *View {
	return &View{
		ID:                 person.ID,
		FullName:           person.FullName,
		PreferredName:      person.PreferredName,
		LanguagePreference: person.LanguagePreference,
	}
}
