package roles

import (
	"lms-backend/internal/model"
)

func GetAllRoles() []model.Role {
	return []model.Role{
		SystemAdmin,
		LibraryAdmin,
		Staff,
		Basic,
	}
}
