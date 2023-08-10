package abilities

import (
	"lms-backend/internal/model"
)

var (
	CanCreatePerson model.Ability = model.Ability{
		Name:        "canCreatePerson",
		Description: "can create person",
	}
	CanUpdatePerson model.Ability = model.Ability{
		Name:        "canUpdatePerson",
		Description: "can update person",
	}
)
