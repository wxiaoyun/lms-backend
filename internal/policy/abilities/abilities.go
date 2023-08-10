package abilities

import (
	"lms-backend/internal/model"
)

var (
	CanManageAll model.Ability = model.Ability{
		Name:        "canManageAll",
		Description: "can manage all",
	}
	CanCreateBook model.Ability = model.Ability{
		Name:        "canCreateBook",
		Description: "can create book",
	}
	CanReadBook model.Ability = model.Ability{
		Name:        "canReadBook",
		Description: "can read book",
	}
	CanUpdateBook model.Ability = model.Ability{
		Name:        "canUpdateBook",
		Description: "can update book",
	}
	CanDeleteBook model.Ability = model.Ability{
		Name:        "canDeleteBook",
		Description: "can delete book",
	}
)
