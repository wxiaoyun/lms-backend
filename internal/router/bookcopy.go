package router

import (
	bookcopyhandler "lms-backend/internal/handler/bookcopy"
	loanhandler "lms-backend/internal/handler/loan"
	reservationhandler "lms-backend/internal/handler/reservation"

	"github.com/gofiber/fiber/v2"
)

func BookcopyRoutes(r fiber.Router) {
	Route(r, "/:bookcopy_id", func(r fiber.Router) {
		r.Delete("/", bookcopyhandler.HandleDelete)
		r.Get("/qrcode", bookcopyhandler.HandleGenerateQRCode)

		Route(r, "/loan", BookLoanRoutes)
		Route(r, "/reservation", BookReservationRoutes)
	})
}

func BookLoanRoutes(r fiber.Router) {
	r.Post("/", loanhandler.HandleLoan)
	r.Patch("/return", loanhandler.HandleReturnByBookcopy)
}

func BookReservationRoutes(r fiber.Router) {
	r.Post("/", reservationhandler.HandleReserve)
}
