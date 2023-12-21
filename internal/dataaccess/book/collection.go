package book

import (
	collection "lms-backend/pkg/collectionquery"
)

func Filters() collection.FilterMap {
	return map[string]collection.Filter{
		"title":     collection.StringLikeFilter("title"),
		"author":    collection.StringLikeFilter("author"),
		"isbn":      collection.StringLikeFilter("isbn"),
		"publisher": collection.StringLikeFilter("publisher"),
		"value":     collection.MultipleColumnStringLikeFilter([]string{"title", "author", "isbn", "publisher"}),
	}
}

func Sorters() collection.SortMap {
	return map[string]collection.Sorter{
		"title":            collection.SortBy("title"),
		"author":           collection.SortBy("author"),
		"isbn":             collection.SortBy("isbn"),
		"publisher":        collection.SortBy("publisher"),
		"publication_date": collection.SortBy("publication_date"),
		"created_at":       collection.SortBy("created_at"),
	}
}
