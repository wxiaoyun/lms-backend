package abilities

import (
	"lms-backend/internal/model"
)

var (
	CanReadUser model.Ability = model.Ability{
		Name:        "canReadUser",
		Description: "can read user",
	}
	CanUpdateUser model.Ability = model.Ability{
		Name:        "canUpdateUser",
		Description: "can update user",
	}
	CanDeleteUser model.Ability = model.Ability{
		Name:        "canDeleteUser",
		Description: "can delete user",
	}
	CanUpdateUserRole model.Ability = model.Ability{
		Name:        "canUpdateRole",
		Description: "can update role",
	}
)
