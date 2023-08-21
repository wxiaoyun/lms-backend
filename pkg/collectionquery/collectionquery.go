package collection

import (
	"github.com/gofiber/fiber/v2"
)

type Order = string
type Query struct {
	Offset  int
	Limit   int
	Search  string
	SortBy  string
	OrderBy string
	Queries map[string]string
}

const (
	offsetKey  = "offset"
	limitKey   = "limit"
	searchKey  = "search"
	sortByKey  = "sortBy"
	orderByKey = "orderBy"

	ASC  Order = "asc"
	DESC Order = "desc"
)

// Returns the collection query sent from the frontend.
// Sets the default value if none is provided
func GetCollectionQueryFromParam(c *fiber.Ctx) *Query {
	offset := c.QueryInt(offsetKey, 0)
	limit := c.QueryInt(limitKey, 25)
	search := c.Query(searchKey, "")

	orderBy := c.Query(orderByKey, DESC)
	if orderBy != ASC && orderBy != DESC {
		// default
		orderBy = DESC
	}
	sortBy := c.Query(sortByKey, "")

	return &Query{
		Offset:  offset,
		Limit:   limit,
		Search:  search,
		SortBy:  sortBy,
		OrderBy: orderBy,
		Queries: c.Queries(),
	}
}
