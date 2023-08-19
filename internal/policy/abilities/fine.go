package abilities

import (
	"lms-backend/internal/model"
)

var (
	CanReadFine model.Ability = model.Ability{
		Name:        "canReadFine",
		Description: "can read fine",
	}
	CanSettleFine model.Ability = model.Ability{
		Name:        "canSettleFine",
		Description: "can settle fine",
	}
	CanDeleteFine model.Ability = model.Ability{
		Name:        "canDeleteFine",
		Description: "can delete fine",
	}
)
