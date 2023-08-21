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

// @Summary List reservations
// @Description List reservations in the library depending on the collection query
// @Tags reservation
// @Accept application/json
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param filter[user_id] query int false "Filter by user ID"
// @Param sortBy query string false "Sort by column name"
// @Param orderBy query string false "Order by direction (asc or desc)"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[[]reservationview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/reservation [get]
func HandleList(c *fiber.Ctx) error {
	err := policy.Authorize(c, readReservationAction, reservationpolicy.ReadPolicy())
	if err != nil {
		return err
	}

	cq := collection.GetCollectionQueryFromParam(c)
	db := database.GetDB()

	totalCount, err := reservation.Count(db)
	if err != nil {
		return err
	}

	dbFiltered := cq.Filter(db, reservation.Filters())

	filteredCount, err := reservation.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered, reservation.Sorters())
	dbPaginated := cq.Paginate(dbSorted)

	rvs, err := reservation.List(dbPaginated)
	if err != nil {
		return err
	}

	var view = []*reservationview.View{}
	for _, r := range rvs {
		//nolint:gosec // loop does not modify struct
		view = append(view, reservationview.ToView(&r))
	}

	return c.JSON(api.Response{
		Data: view,
		Meta: api.Meta{
			TotalCount:    totalCount,
			FilteredCount: filteredCount,
		},
		Messages: api.Messages(
			api.SilentMessage("reservations listed successfully"),
		),
	})
}
