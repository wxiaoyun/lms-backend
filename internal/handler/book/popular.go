package bookhandler

import (
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"
	"lms-backend/internal/view/bookview"

	"github.com/gofiber/fiber/v2"
)

func HandlePopular(c *fiber.Ctx) error {
	err := policy.Authorize(c, readBookAction, bookpolicy.ListPolicy())
	if err != nil {
		return err
	}

	db := database.GetDB()
	blc, err := book.ListPopularBooks(db)
	if err != nil {
		return err
	}

	var view = []bookview.PopularView{}
	for _, b := range blc {
		//nolint:gosec // loop does not modify struct
		view = append(view, *bookview.ToPopularView(&b))
	}

	return c.JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SilentMessage("Popular books listed successfully"),
		),
	})
}
