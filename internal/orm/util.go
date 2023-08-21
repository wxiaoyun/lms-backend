package orm

import (
	"lms-backend/util/sliceutil"

	"gorm.io/gorm"
)

func Join(joinQuery string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB { return db.Joins(joinQuery) }
}

func JoinAll(joinQueries []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return sliceutil.Reduce(
			joinQueries,
			func(db *gorm.DB, query string) *gorm.DB { return db.Scopes(Join(query)) },
			db,
		)
	}
}
