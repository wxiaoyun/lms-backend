package collection

import (
	"gorm.io/gorm"
)

// Paginates the database query based on the collection query
//
// To be called after the Filter() and Sort() methods.
func (q *Query) Paginate(db *gorm.DB) *gorm.DB {
	return db.Scopes(func(*gorm.DB) *gorm.DB {
		return db.Offset(q.Offset).Limit(q.Limit)
	})
}
