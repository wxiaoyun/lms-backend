package collection

import (
	"fmt"
	"lms-backend/internal/orm"
	"lms-backend/util/sliceutil"
	"regexp"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

var (
	FilterRegex      = regexp.MustCompile(`filter\[[^\]]*\]`)
	FilterValueRegex = regexp.MustCompile(`filter\[([^\]]*)\]`)
)

type Filter func(string) func(*gorm.DB) *gorm.DB

// FilterMap is a map of filter keys to filter functions
//
// Key should the the filter key (i.e. "filter[<key>]")
type FilterMap map[string]Filter

// Filters the database query based on the collection query
//
// The filters map should be a map of filter keys (i.e. "filter[<key>]") to
// filter functions. The filter functions should take a string as an argument
// and return a function that takes a *gorm.DB and returns a *gorm.DB.

// Example:
//
//	filters := map[string]Filter{
//
//		"age": (x) -> db.Where("age = ?", x),
//
//		"name": (x) -> db.Where("name ILIKE ?", x),
//
//	}
//
// joinQueries is a list of join queries to be applied to the database query via the db.Joins() method.
//
// Example:
//
//	joinQueries := []string{
//
//		"JOIN users ON users.id = posts.user_id",
//
//	}
func (q *Query) Filter(db *gorm.DB, filters FilterMap) *gorm.DB {
	// Key should be the column name and value should be the value to filter by
	for key, value := range q.Queries {
		// Test if the key is a filter key (filter[...])
		if !FilterRegex.MatchString(key) {
			continue
		}

		matches := FilterValueRegex.FindStringSubmatch(key)

		// If there are no matches or the first match is empty, skip this key
		if len(matches) < 1 || matches[1] == "" {
			continue
		}

		// The first match is the whole string, the second match is the filter key
		filterKey := matches[1]
		if filter, ok := filters[filterKey]; ok {
			db = db.Scopes(filter(value))
		}
	}

	return db
}

func genericFilter(columnName, operator string, joinQueries ...string) Filter {
	return func(value string) func(db *gorm.DB) *gorm.DB {
		return func(db *gorm.DB) *gorm.DB {
			return db.Scopes(orm.JoinAllIfNotJoined(joinQueries)).Where(
				fmt.Sprintf("%s %s ?", columnName, operator),
				value,
			)
		}
	}
}

func StringLikeFilter(columnName string, joinQueries ...string) Filter {
	return func(value string) func(db *gorm.DB) *gorm.DB {
		return func(db *gorm.DB) *gorm.DB {
			return db.Scopes(orm.JoinAllIfNotJoined(joinQueries)).Where(
				fmt.Sprintf("%s ILIKE ?", columnName),
				fmt.Sprintf("%%%s%%", value),
			)
		}
	}
}

func StringEqualFilter(columnName string, joinQueries ...string) Filter {
	return genericFilter(columnName, "=", joinQueries...)
}

// When the value is a comma-separated list of values, this filter will return
// a function that will filter the database query by the column name being equal
// to any of the values in the list.
//
// One column, multiple values
func MultipleStringEqualFilter(columnName string, joinQueries ...string) Filter {
	return func(value string) func(db *gorm.DB) *gorm.DB {
		// Split the value by commas and surround each value with single quotes
		values := sliceutil.Map(strings.Split(value, ","), func(s string) string {
			return fmt.Sprintf("'%s'", s)
		})

		return func(db *gorm.DB) *gorm.DB {
			return db.Scopes(orm.JoinAllIfNotJoined(joinQueries)).Where(
				fmt.Sprintf("%s IN ?", columnName),
				fmt.Sprintf("(%s)", strings.Join(values, ",")),
			)
		}
	}
}

// When the value is a comma-separated list of values, this filter will return
// a function that will filter the database query by the column name being similar
// to any of the values in the list.
//
// One column, multiple values
func MultipleStringLikeFilter(columnName string, joinQueries ...string) Filter {
	return func(value string) func(db *gorm.DB) *gorm.DB {
		// Split the value by commas and convert them to ILIKE conditions (i.e. "column ILIKE '%value%'")
		conditions := sliceutil.Map(strings.Split(value, ","), func(s string) string {
			return fmt.Sprintf("%s ILIKE %%%s%%", columnName, s)
		})

		return func(db *gorm.DB) *gorm.DB {
			return db.Scopes(orm.JoinAllIfNotJoined(joinQueries)).Where(
				fmt.Sprintf("(%s)", strings.Join(conditions, " OR ")),
			)
		}
	}
}

// When the value is a single value, this filter will return
// a function that will filter the database query if the column name is similar to the value.
//
// Multiple columns, one value
func MultipleColumnStringLikeFilter(columnNames []string, joinQueries ...string) Filter {
	return func(value string) func(db *gorm.DB) *gorm.DB {
		// Split the value by commas and convert them to ILIKE conditions (i.e. "column ILIKE '%value%'")
		conditions := sliceutil.Map(columnNames, func(s string) string {
			return fmt.Sprintf("%s ILIKE '%%%s%%'", s, value)
		})

		return func(db *gorm.DB) *gorm.DB {
			return db.Scopes(orm.JoinAllIfNotJoined(joinQueries)).Where(
				fmt.Sprintf("(%s)", strings.Join(conditions, " OR ")),
			)
		}
	}
}

func GenericBoolFilter(value bool, columnName string, operator string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(
			fmt.Sprintf("%s %s ?", columnName, operator),
			value,
		)
	}
}

func IntEqualFilter(columnName string, joinQueries ...string) Filter {
	return genericFilter(columnName, "=", joinQueries...)
}

func IntGreaterThanOrEqualFilter(columnName string, joinQueries ...string) Filter {
	return genericFilter(columnName, ">=", joinQueries...)
}

func IntLessThanOrEqualFilter(columnName string, joinQueries ...string) Filter {
	return genericFilter(columnName, "<=", joinQueries...)
}

func IntGreaterThanFilter(columnName string, joinQueries ...string) Filter {
	return genericFilter(columnName, ">", joinQueries...)
}

func IntLessThanFilter(columnName string, joinQueries ...string) Filter {
	return genericFilter(columnName, "<", joinQueries...)
}

// When the value is a comma-separated list of values, this filter will return
// a function that will filter the database query by the column name equal
// to any of the values in the list.
//
// One column, multiple values
func MultipleIntEqualFilter(columnName string, joinQueries ...string) Filter {
	return func(value string) func(db *gorm.DB) *gorm.DB {
		// Split the value by commas and convert them to ILIKE conditions (i.e. "column ILIKE '%value%'")
		conditions := sliceutil.Map(strings.Split(value, ","), func(s string) string {
			val, err := strconv.Atoi(s)
			if err != nil {
				return "1 = 0" // Return a condition that will never be true
			}
			return fmt.Sprintf("%s = %d", columnName, val)
		})

		return func(db *gorm.DB) *gorm.DB {
			return db.Scopes(orm.JoinAllIfNotJoined(joinQueries)).Where(
				fmt.Sprintf("(%s)", strings.Join(conditions, " OR ")),
			)
		}
	}
}
