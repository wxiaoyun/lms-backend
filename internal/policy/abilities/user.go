package abilities

import (
	"lms-backend/internal/model"
)

var (
	CanUpdateUser model.Ability = model.Ability{
		Name:        "canUpdateUser",
		Description: "can update user",
	}
	CanUpdateRole model.Ability = model.Ability{
		Name:        "canUpdateRole",
		Description: "can update role",
	}
)
