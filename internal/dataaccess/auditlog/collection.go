package auditlog

import (
	collection "lms-backend/pkg/collectionquery"
)

func Filters() collection.FilterMap {
	return map[string]collection.Filter{
		"username": collection.StringLikeFilter("users.username", JoinUser),
		"action":   collection.StringLikeFilter("action"),
		"value":    collection.MultipleColumnStringLikeFilter([]string{"action", "users.username"}, JoinUser),
	}
}

func Sorters() collection.SortMap {
	return map[string]collection.Sorter{
		"action":     collection.SortBy("action"),
		"created_at": collection.SortBy("created_at"),
	}
}
