package router

import (
	"lms-backend/internal/config"
	"lms-backend/internal/handler/auth"
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
	r.Post("/user", userhandler.HandleCreate)

	Route(r, "book", func(r fiber.Router) {
		r.Get("/", bookhandler.HandleList)
		r.Get("/popular", middleware.CacheMiddleware(middleware.VLongExp), bookhandler.HandlePopular)
		r.Get("/:book_id", bookhandler.HandleRead)
	})
}

func PrivateRoutes(r fiber.Router) {
	r.Get("/auth/signout", auth.HandleSignOut)
	Route(r, "/user", UserRoutes)
	Route(r, "/book", BookRoutes)
	Route(r, "/bookcopy", BookcopyRoutes)
	Route(r, "/bookmark", BookmarkRoutes)
	Route(r, "/loan", LoanRoutes)
	Route(r, "/reservation", ReservationRoutes)
	Route(r, "/fine", FineRoutes)
	Route(r, "/audit_log", AuditLogRoutes)
	Route(r, "/external", ExternalRoutes)
}
