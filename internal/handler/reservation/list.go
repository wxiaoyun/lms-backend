package reservationhandler

import (
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/reservation"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/reservationpolicy"
	"lms-backend/internal/view/reservationview"
	collection "lms-backend/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
)

const (
	listBookReservationAction = "list book on reservation"
)

func HandleList(c *fiber.Ctx) error {
	err := policy.Authorize(c, listBookReservationAction, reservationpolicy.ReadBookPolicy())
	if err != nil {
		return err
	}

	cq := collection.GetCollectionQueryFromParam(c)
	db := database.GetDB()

	totalCount, err := reservation.Count(db)
	if err != nil {
		return err
	}

	dbFiltered := cq.Filter(db, reservation.Filters(), reservation.JoinBookCopy, reservation.JoinBook)

	filteredCount, err := reservation.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered, reservation.Sorters())
	dbPaginated := cq.Paginate(dbSorted)

	res, err := reservation.ListWithBookUser(dbPaginated)
	if err != nil {
		return err
	}

	var view = []reservationview.DetailedView{}
	for _, r := range res {
		//nolint:gosec // loop does not modify struct
		view = append(view, *reservationview.ToDetailedView(&r))
	}

	return c.JSON(api.Response{
		Data: view,
		Meta: api.Meta{
			TotalCount:    totalCount,
			FilteredCount: filteredCount,
		},
		Messages: api.Messages(
			api.SilentMessage("books listed successfully"),
		),
	})
}
