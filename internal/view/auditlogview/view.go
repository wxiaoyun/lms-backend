package auditlogview

import (
	"technical-test/internal/model"
)

type View struct {
	ID     uint   `json:"id,omitempty"`
	Action string `json:"action"`
	UserID uint   `json:"user_id"`
}

func ToView(auditLog *model.AuditLog) *View {
	return &View{
		ID:     auditLog.ID,
		Action: auditLog.Action,
		UserID: auditLog.UserID,
	}
}
