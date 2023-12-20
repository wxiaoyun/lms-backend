package reservationhandler

import (
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/reservationpolicy"
	"lms-backend/internal/view/bookview"
	collection "lms-backend/pkg/collectionquery"
	"lms-backend/pkg/error/internalerror"

	"github.com/gofiber/fiber/v2"
)

const (
	listBookReservationAction = "list book on reservation"
)

// @Summary List books on reservation
// @Description Lists books on reservation
// @Tags loan
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[bookview.BookReservationView]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/reservation/book [get]
func HandleListBook(c *fiber.Ctx) error {
	err := policy.Authorize(c, listBookReservationAction, reservationpolicy.ReadBookPolicy())
	if err != nil {
		return err
	}

	cq := collection.GetCollectionQueryFromParam(c)
	db := database.GetDB()

	totalCount, err := book.Count(db)
	if err != nil {
		return err
	}

	dbFiltered := cq.Filter(db, book.ReservationFilters())

	filteredCount, err := book.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered, book.Sorters())
	dbPaginated := cq.Paginate(dbSorted)

	books, err := book.ListWithReservation(dbPaginated)
	if err != nil {
		return err
	}

	var view = []*bookview.BookReservationView{}
	for _, b := range books {
		if (b.Reservations == nil) || (len(b.Reservations) == 0) {
			return internalerror.InternalServerError("book does not have any reservation")
		}
		//nolint:gosec // loop does not modify struct
		view = append(view, bookview.ToBookReservationView(&b, &b.Reservations[0]))
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
