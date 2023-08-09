package audit

import (
	"technical-test/internal/model"
	"technical-test/internal/session"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Starts a transaction and returns a function that should be deferred.
//
// The function will commit the transaction if no panic occurs, otherwise it will rollback.
//
// The function will also create an audit log entry with the provided action message.
func Begin(c *fiber.Ctx, db *gorm.DB, action string) (*gorm.DB, func()) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		// Default to system admin
		userID = 1
	}

	var deferedRollBackOrCommit = func() {
		//nolint
		if r := recover(); r != nil {
			db.Rollback()
			return
		}

		auditLog := model.AuditLog{
			UserID: uint(userID),
			Action: action,
		}

		if err := auditLog.Create(db); err != nil {
			db.Rollback()
			return
		}

		db.Commit()
	}

	return db.Begin(), deferedRollBackOrCommit
}
