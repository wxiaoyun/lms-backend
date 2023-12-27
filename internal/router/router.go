package router

import (
	"lms-backend/internal/config"
	bookhandler "lms-backend/internal/handler/book"
	userhandler "lms-backend/internal/handler/user"
	"lms-backend/internal/middleware"
	sessionmiddleware "lms-backend/internal/middleware/session"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, cfg *config.Config) {
	session.SetupStore()
	middleware.SetupCors(app, cfg)
	middleware.SetupRecover(app)
	middleware.SetupLogger(app)

	v1Routes := app.Group("/v1")

	publicRoutes := v1Routes.Group("/")
	Route(publicRoutes, "/", PublicRoutes)

	privateRoutes := v1Routes.Group("/", sessionmiddleware.SessionMiddleware)
	Route(privateRoutes, "/", PrivateRoutes)
}

func PublicRoutes(r fiber.Router) {
	Route(r, "/health", HealthRoutes)
	Route(r, "/auth", AuthRoutes)
	r.Get("/current", userhandler.HandleGetCurrentUser)
	r.Get("/book", middleware.CacheMiddleware(middleware.VLongExp), bookhandler.HandleList)
}

func PrivateRoutes(r fiber.Router) {
	Route(r, "/user", UserRoutes)
	Route(r, "/book", BookRoutes)
	Route(r, "/bookmark", BookmarkRoutes)
	Route(r, "/loan", LoanRoutes)
	Route(r, "/reservation", ReservationRoutes)
	Route(r, "/fine", FineRoutes)
	Route(r, "/audit_log", AuditLogRoutes)
}
