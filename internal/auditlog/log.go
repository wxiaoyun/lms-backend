package audit

import (
	"lms-backend/internal/model"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Starts a transaction and returns a function that should be deferred.
//
// The function will commit the transaction if no panic occurs, otherwise it will rollback.
//
// The function will also create an audit log entry with the provided action message.
func Begin(c *fiber.Ctx, db *gorm.DB, action string) (*gorm.DB, func(error)) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		// Default to system admin
		userID = 1
	}

	tx := db.Begin()

	var deferedRollBackOrCommit = func(err error) {
		//nolint
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
			return
		}

		auditLog := model.AuditLog{
			UserID: uint(userID),
			Action: action,
		}

		if err := auditLog.Create(tx); err != nil {
			tx.Rollback()
			return
		}

		tx.Commit()
	}

	return tx, deferedRollBackOrCommit
}
