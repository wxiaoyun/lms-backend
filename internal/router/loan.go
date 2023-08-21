package router

import (
	loanhandler "lms-backend/internal/handler/loan"

	"github.com/gofiber/fiber/v2"
)

func LoanRoutes(r fiber.Router) {
	r.Get("/", loanhandler.HandleList)
}
