package router

import (
	bookhandler "lms-backend/internal/handler/book"
	loanhandler "lms-backend/internal/handler/loan"
	reservationhandler "lms-backend/internal/handler/reservation"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(r fiber.Router) {
	r.Post("/", bookhandler.HandleCreate)
	r.Get("/", bookhandler.HandleList)

	Route(r, "/:id", func(r fiber.Router) {
		r.Get("/", bookhandler.HandleRead)
		r.Patch("/", bookhandler.HandleUpdate)
		r.Delete("/", bookhandler.HandleDelete)

		Route(r, "/loan", LoanRoutes)
		Route(r, "/reservation", ReservationRoutes)
	})
}

func LoanRoutes(r fiber.Router) {
	r.Post("/", loanhandler.HandleLoan)

	Route(r, "/:loan_id", func(r fiber.Router) {
		r.Get("/", loanhandler.HandleRead)
		r.Delete("/", loanhandler.HandleDelete)
		r.Patch("/return", loanhandler.HandleReturn)
		r.Patch("/renew", loanhandler.HandleRenew)
	})
}

func ReservationRoutes(r fiber.Router) {
	r.Post("/", reservationhandler.HandleReserve)

	Route(r, "/:reservation_id", func(r fiber.Router) {
		r.Get("/", reservationhandler.HandleRead)
		r.Delete("/", reservationhandler.HandleDelete)
		r.Patch("/cancel", reservationhandler.HandleCancel)
		r.Patch("/checkout", reservationhandler.HandleCheckout)
	})
}
