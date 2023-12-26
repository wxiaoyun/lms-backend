package router

import (
	bookhandler "lms-backend/internal/handler/book"
	loanhandler "lms-backend/internal/handler/loan"
	reservationhandler "lms-backend/internal/handler/reservation"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(r fiber.Router) {
	r.Post("/", bookhandler.HandleCreate)

	Route(r, "/:book_id", func(r fiber.Router) {
		r.Get("/", bookhandler.HandleRead)
		r.Patch("/", bookhandler.HandleUpdate)
		r.Delete("/", bookhandler.HandleDelete)

		Route(r, "/loan", BookLoanRoutes)
		Route(r, "/reservation", BookReservationRoutes)
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
