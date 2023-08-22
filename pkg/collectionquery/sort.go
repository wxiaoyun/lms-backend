package collection

import (
	"fmt"
	"lms-backend/internal/orm"

	"gorm.io/gorm"
)

type Sorter func(string) func(*gorm.DB) *gorm.DB

type SortMap map[string]Sorter

// Sorts the database query based on the collection query
//
// To be called after the Filter() method.
func (q *Query) Sort(db *gorm.DB, sorters SortMap, joinQueries ...string) *gorm.DB {
	if q.sortBy == "" {
		return db
	}

	db = db.Scopes(orm.JoinAll(joinQueries))

	sorter, ok := sorters[q.sortBy]
	if !ok {
		return db
	}

	return db.Scopes(sorter(q.orderBy))
}

func SortBy(columnName string) Sorter {
	return func(order string) func(db *gorm.DB) *gorm.DB {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order(fmt.Sprintf("%s %s", columnName, order))
		}
	}
}
