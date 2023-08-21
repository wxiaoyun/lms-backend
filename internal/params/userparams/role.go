package userparams

import (
	"lms-backend/pkg/error/externalerrors"
)

type UpdateRoleParams struct {
	RoleID int64 `json:"role_id"`
}

func (p *UpdateRoleParams) Validate() error {
	if p.RoleID == 0 {
		return externalerrors.BadRequest("Role ID is required.")
	}

	return nil
}
