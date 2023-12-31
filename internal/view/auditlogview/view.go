package auditlogview

import (
	"lms-backend/internal/model"
	"time"
)

type View struct {
	ID     uint   `json:"id,omitempty"`
	Action string `json:"action"`
	UserID uint   `json:"user_id"`
	Date   string `json:"date"`
}

func ToView(auditLog *model.AuditLog) *View {
	return &View{
		ID:     auditLog.ID,
		Action: auditLog.Action,
		UserID: auditLog.UserID,
		Date:   auditLog.Date.Format(time.RFC3339),
	}
}
