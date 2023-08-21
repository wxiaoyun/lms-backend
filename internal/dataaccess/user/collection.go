package user

import (
	collection "lms-backend/pkg/collectionquery"
)

func Filters() collection.FilterMap {
	return map[string]collection.Filter{
		"username": collection.StringLikeFilter("username"),
		"email":    collection.StringLikeFilter("email"),
	}
}

func Sorters() collection.SortMap {
	return map[string]collection.Sorter{
		"username":           collection.SortBy("username"),
		"email":              collection.SortBy("email"),
		"sign_in_count":      collection.SortBy("sign_in_count"),
		"current_sign_in_at": collection.SortBy("current_sign_in_at"),
		"created_at":         collection.SortBy("created_at"),
	}
}
