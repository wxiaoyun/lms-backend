package roles

import (
	"lms-backend/internal/model"
)

var (
	SystemAdmin model.Role = model.Role{
		Name: "System Admin",
	}
	LibraryAdmin model.Role = model.Role{
		Name: "Library Admin",
	}
	Staff model.Role = model.Role{
		Name: "Staff",
	}
	Basic model.Role = model.Role{
		Name: "Basic",
	}
)
