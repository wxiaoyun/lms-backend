package fine

import (
	collection "lms-backend/pkg/collectionquery"
)

func Filters() collection.FilterMap {
	return map[string]collection.Filter{
		"status":         collection.StringEqualFilter("fines.status"),
		"user_id":        collection.MultipleIntEqualFilter("fines.user_id"),
		"loan_id":        collection.MultipleIntEqualFilter("loan_id"),
		"users.username": collection.StringLikeFilter("users.username", JoinUser),
		"books.value":    collection.MultipleColumnStringLikeFilter([]string{"books.title", "books.author", "books.isbn", "books.publisher"}, JoinLoan, JoinBookCopy, JoinBook),
		"value":          collection.MultipleColumnStringLikeFilter([]string{"books.title", "books.author", "books.isbn", "books.publisher", "users.username"}, JoinLoan, JoinBookCopy, JoinBook, JoinUser),
	}
}

func Sorters() collection.SortMap {
	return map[string]collection.Sorter{
		"amount":     collection.SortBy("amount"),
		"created_at": collection.SortBy("fines.created_at"),
	}
}
