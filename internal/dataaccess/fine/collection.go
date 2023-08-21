package fine

import (
	collection "lms-backend/pkg/collectionquery"
)

func Filters() collection.FilterMap {
	return map[string]collection.Filter{
		"status":  collection.StringEqualFilter("status"),
		"user_id": collection.MultipleIntEqualFilter("user_id"),
		"loan_id": collection.MultipleIntEqualFilter("loan_id"),
	}
}

func Sorters() collection.SortMap {
	return map[string]collection.Sorter{
		"amount":     collection.SortBy("amount"),
		"created_at": collection.SortBy("created_at"),
	}
}
