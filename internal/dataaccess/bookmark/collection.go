package bookmark

import (
	collection "lms-backend/pkg/collectionquery"
)

func Filters() collection.FilterMap {
	return map[string]collection.Filter{
		"book_id": collection.IntEqualFilter("book_id"),
		"user_id": collection.IntEqualFilter("user_id"),
	}
}

func Sorters() collection.SortMap {
	return map[string]collection.Sorter{
		"created_at": collection.SortBy("created_at"),
	}
}
