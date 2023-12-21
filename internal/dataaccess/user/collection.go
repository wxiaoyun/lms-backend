package user

import (
	collection "lms-backend/pkg/collectionquery"
)

func Filters() collection.FilterMap {
	return map[string]collection.Filter{
		"username":        collection.StringLikeFilter("username"),
		"email":           collection.StringLikeFilter("email"),
		"full_name":       collection.StringLikeFilter("people.full_name", JoinPerson),
		"Preferred_named": collection.StringLikeFilter("people.preferred_name", JoinPerson),
		"value":           collection.MultipleColumnStringLikeFilter([]string{"username", "email", "people.full_name", "people.preferred_name"}, JoinPerson),
	}
}

func Sorters() collection.SortMap {
	return map[string]collection.Sorter{
		"username":       collection.SortBy("username"),
		"email":          collection.SortBy("email"),
		"created_at":     collection.SortBy("created_at"),
		"full_name":      collection.SortBy("people.full_name", JoinPerson),
		"preferred_name": collection.SortBy("people.preferred_name", JoinPerson),
	}
}
