package auditlogview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/userview"
)

type DetailedView struct {
	View
	User userview.View `json:"user"`
}

func ToDetailedView(a *model.AuditLog) *DetailedView {
	return &DetailedView{
		View: *ToView(a),
		User: *userview.ToView(a.User),
	}
}
