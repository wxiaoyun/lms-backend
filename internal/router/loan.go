package router

import (
	loanhandler "lms-backend/internal/handler/loan"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func LoanRoutes(r fiber.Router) {
	r.Get("/", middleware.CacheMiddleware(middleware.ShortExp), loanhandler.HandleList)
	r.Post("/", loanhandler.HandleCreate)

	Route(r, "/:loan_id", func(r fiber.Router) {
		r.Get("/", loanhandler.HandleRead)
		r.Delete("/", loanhandler.HandleDelete)
		r.Patch("/return", loanhandler.HandleReturn)
		r.Patch("/renew", loanhandler.HandleRenew)
	})
}
