package router

import (
	reservationhandler "lms-backend/internal/handler/reservation"

	"github.com/gofiber/fiber/v2"
)

func ReservationRoutes(r fiber.Router) {
	r.Get("/", reservationhandler.HandleList)
	r.Post("/", reservationhandler.HandleCreate)
	r.Post("/book", reservationhandler.HandleCreateByBook)

	Route(r, "/:reservation_id", func(r fiber.Router) {
		r.Get("/", reservationhandler.HandleRead)
		r.Patch("/cancel", reservationhandler.HandleCancel)
	})
}
