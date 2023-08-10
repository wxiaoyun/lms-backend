package abilities

import (
	"lms-backend/internal/model"
)

var (
	CanReadAuditLog model.Ability = model.Ability{
		Name:        "canReadAuditLog",
		Description: "can read audit log",
	}
	CanCreateAuditLog model.Ability = model.Ability{
		Name:        "canCreateAuditLog",
		Description: "can create audit log",
	}
)
