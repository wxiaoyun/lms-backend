package loan

import (
	collection "lms-backend/pkg/collectionquery"
)

func Filters() collection.FilterMap {
	return map[string]collection.Filter{
		"status":  collection.StringEqualFilter("status"),
		"user_id": collection.MultipleIntEqualFilter("user_id"),
		"book_id": collection.MultipleIntEqualFilter("book_id"),
	}
}

func Sorters() collection.SortMap {
	return map[string]collection.Sorter{
		"borrowed_date": collection.SortBy("borrowed_date"),
		"due_date":      collection.SortBy("due_date"),
		"created_at":    collection.SortBy("created_at"),
	}
}
