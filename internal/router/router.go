package router

import (
	"lms-backend/internal/config"
	userhandler "lms-backend/internal/handler/user"
	"lms-backend/internal/middleware"
	sessionmiddleware "lms-backend/internal/middleware/session"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, cfg *config.Config) {
	middleware.SetupCors(app, cfg)
	middleware.SetupRecover(app)
	middleware.SetupRecover(app)
	session.SetupStore()

	v1Routes := app.Group("/api/v1")

	publicRoutes := v1Routes.Group("/")
	Route(publicRoutes, "/", PublicRoutes)

	privateRoutes := v1Routes.Group("/")
	privateRoutes.Use(sessionmiddleware.SessionMiddleware)
	Route(privateRoutes, "/", PrivateRoutes)
}

func PublicRoutes(r fiber.Router) {
	Route(r, "/health", HealthRoutes)
	Route(r, "/auth", AuthRoutes)
	r.Get("/current", userhandler.HandleGetCurrentUser)
}

func PrivateRoutes(r fiber.Router) {
	Route(r, "/user", UserRoutes)
	Route(r, "/audit_log", AuditLogRoutes)
	Route(r, "/book", BookRoutes)
	Route(r, "/loan", LoanRoutes)
	Route(r, "/reservation", ReservationRoutes)
	Route(r, "/fine", FineRoutes)
}
