package router

import (
	reservationhandler "lms-backend/internal/handler/reservation"

	"github.com/gofiber/fiber/v2"
)

func ReservationRoutes(r fiber.Router) {
	r.Get("/", reservationhandler.HandleList)
	r.Post("/", reservationhandler.HandleCreate)

	Route(r, "/:reservation_id", func(r fiber.Router) {
		r.Get("/", reservationhandler.HandleRead)
		r.Delete("/", reservationhandler.HandleDelete)
		r.Patch("/cancel", reservationhandler.HandleCancel)
		r.Patch("/checkout", reservationhandler.HandleCheckout)
	})
}
