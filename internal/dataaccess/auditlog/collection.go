package auditlog

import (
	collection "lms-backend/pkg/collectionquery"
)

func Filters() map[string]collection.Filter {
	return map[string]collection.Filter{
		"action": collection.StringLikeFilter("action"),
	}
}
