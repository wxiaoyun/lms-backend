package reservation

import (
	collection "lms-backend/pkg/collectionquery"
)

func Filters() collection.FilterMap {
	return map[string]collection.Filter{
		"status":         collection.StringEqualFilter("reservations.status"),
		"user_id":        collection.MultipleIntEqualFilter("reservations.user_id"),
		"book_id":        collection.MultipleIntEqualFilter("reservations.book_id"),
		"users.username": collection.StringLikeFilter("users.username", JoinUser),
		"books.value":    collection.MultipleColumnStringLikeFilter([]string{"books.title", "books.author", "books.isbn", "books.publisher"}, JoinBookCopy, JoinBook),
		"value":          collection.MultipleColumnStringLikeFilter([]string{"books.title", "books.author", "books.isbn", "books.publisher", "users.username"}, JoinBookCopy, JoinBook, JoinUser),
	}
}

func Sorters() collection.SortMap {
	return map[string]collection.Sorter{
		"reservation_date": collection.SortBy("reservation_date"),
		"created_at":       collection.SortBy("created_at"),
	}
}
