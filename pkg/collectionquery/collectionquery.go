package collection

import (
	"github.com/gofiber/fiber/v2"
)

type Order = string
type Query struct {
	offset  int
	limit   int
	sortBy  string
	orderBy string
	Queries map[string]string
}

const (
	offsetKey  = "offset"
	limitKey   = "limit"
	sortByKey  = "sortBy"
	orderByKey = "orderBy"

	ASC  Order = "asc"
	DESC Order = "desc"
)

// Returns the collection query sent from the frontend.
// Sets the default value if none is provided
func GetCollectionQueryFromParam(c *fiber.Ctx) *Query {
	offset := c.QueryInt(offsetKey, 0)
	limit := c.QueryInt(limitKey, 10)

	orderBy := c.Query(orderByKey, DESC)
	if orderBy != ASC && orderBy != DESC {
		// default
		orderBy = DESC
	}
	sortBy := c.Query(sortByKey, "")

	return &Query{
		offset:  offset,
		limit:   limit,
		sortBy:  sortBy,
		orderBy: orderBy,
		Queries: c.Queries(),
	}
}
