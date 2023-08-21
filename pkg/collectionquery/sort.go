package collection

import (
	"fmt"
	"lms-backend/internal/orm"

	"gorm.io/gorm"
)

type SortBy func(string, string) func(*gorm.DB) *gorm.DB

// Sorts the database query based on the collection query
//
// To be called after the Filter() method.
func (q *Query) Sort(db *gorm.DB, joinQueries ...string) *gorm.DB {
	if q.SortBy == "" {
		return db
	}

	db = db.Scopes(orm.JoinAll(joinQueries))

	return db.Scopes(func(*gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", q.SortBy, q.OrderBy))
	})
}

// Paginates the database query based on the collection query
//
// To be called after the Filter() and Sort() methods.
func (q *Query) Paginate(db *gorm.DB) *gorm.DB {
	return db.Scopes(func(*gorm.DB) *gorm.DB {
		return db.Offset(q.Offset).Limit(q.Limit)
	})
}
