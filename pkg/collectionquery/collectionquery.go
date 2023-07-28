package collection

import (
	"github.com/gofiber/fiber/v2"
)

type Query struct {
	Offset int
	Limit  int
	Search string
}

const (
	offsetKey = "offset"
	limitKey  = "limit"
	searchKey = "search"
)

// Returns the collection query sent from the frontend.
// Sets the default value if none is provided
func GetCollectionQueryFromParam(c *fiber.Ctx) *Query {
	offset := c.QueryInt(offsetKey, 0)
	limit := c.QueryInt(limitKey, 25)
	search := c.Query(searchKey, "")
	return &Query{
		Offset: offset,
		Limit:  limit,
		Search: search,
	}
}
