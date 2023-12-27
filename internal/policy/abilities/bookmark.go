package abilities

import (
	"lms-backend/internal/model"
)

var (
	CanReadBookMark model.Ability = model.Ability{
		Name:        "canReadBookMark",
		Description: "can read book mark",
	}
	CanCreateBookMark model.Ability = model.Ability{
		Name:        "canCreateBookMark",
		Description: "can create book mark",
	}
	CanDeleteBookMark model.Ability = model.Ability{
		Name:        "canDeleteBookMark",
		Description: "can delete book mark",
	}
)
