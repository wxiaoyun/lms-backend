package orm

import (
	"gorm.io/gorm"
)

// NewSession creates a fresh session without any prior queries.
func NewSession(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{NewDB: true})
}
