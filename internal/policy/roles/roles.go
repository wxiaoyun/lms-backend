package roles

import (
	"lms-backend/internal/model"
)

// IMPORTANT:
// The order of the roles in this slice is important.
// The first role is the highest rank. The last role is the lowest rank.
// Higher rank means lower ID number.
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
