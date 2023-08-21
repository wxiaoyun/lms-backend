package orm

import (
	"gorm.io/gorm"
)

// CloneSession creates a new session for the given database connection.
func CloneSession(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{})
}
