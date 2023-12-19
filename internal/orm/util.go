package orm

import (
	"lms-backend/util/sliceutil"

	"gorm.io/gorm"
)

func Join(joinQuery string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB { return db.Joins(joinQuery) }
}

func containsJoin(db *gorm.DB, query string) bool {
	for _, join := range db.Statement.Joins {
		if join.Name == query {
			return true
		}
	}
	return false
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

func JoinIfNotJoined(query string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == "" || containsJoin(db, query) {
			return db
		}
		return db.Joins(query)
	}
}

func JoinAllIfNotJoined(queries []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return sliceutil.Reduce(
			queries,
			func(db *gorm.DB, query string) *gorm.DB { return db.Scopes(JoinIfNotJoined(query)) },
			db,
		)
	}
}

func EmptyFilter(_ string) func(*gorm.DB) *gorm.DB { return func(db *gorm.DB) *gorm.DB { return db } }
