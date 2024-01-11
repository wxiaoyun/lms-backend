package router

import (
	bookhandler "lms-backend/internal/handler/book"
	bookcopyhandler "lms-backend/internal/handler/bookcopy"
	bookmarkhandler "lms-backend/internal/handler/bookmark"
	loanhandler "lms-backend/internal/handler/loan"
	reservationhandler "lms-backend/internal/handler/reservation"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(r fiber.Router) {
	r.Post("/", bookhandler.HandleCreate)

	Route(r, "/:book_id", func(r fiber.Router) {
		r.Patch("/", bookhandler.HandleUpdate)
		r.Delete("/", bookhandler.HandleDelete)

		Route(r, "/loan", BookLoanRoutes)
		Route(r, "/reservation", BookReservationRoutes)
		Route(r, "/bookmark", BookBookmarkRoutes)
		Route(r, "/bookcopy", BookBookcopyRoutes)
	})

	Route(r, "/autocomplete", func(r fiber.Router) {
		r.Get("/:value", middleware.CacheMiddleware(middleware.ShortExp), bookhandler.HandleAutoComplete)
	})
}

func BookLoanRoutes(r fiber.Router) {
	r.Post("/", loanhandler.HandleLoan)
}

func BookReservationRoutes(r fiber.Router) {
	r.Post("/", reservationhandler.HandleReserve)
}

func BookBookmarkRoutes(r fiber.Router) {
	r.Post("/", bookmarkhandler.HandleCreate)
}

func BookBookcopyRoutes(r fiber.Router) {
	r.Post("/", bookcopyhandler.HandleCreate)
}
