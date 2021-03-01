package routes

import (
	"github.com/barandemirbas/go-jwt-server/api"
	"github.com/barandemirbas/go-jwt-server/middlewares"
	"github.com/barandemirbas/go-jwt-server/security"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(routes *fiber.App) {
	routes.Post("/register", security.Register)
	routes.Post("/login", security.Login)
	routes.Get("/api/v1/book/:id?", api.GetBook)
	routes.Post("/api/v1/book", middlewares.IsAuthorized(), api.AddBook)
	routes.Put("/api/v1/book/:id", middlewares.IsAuthorized(), api.UpdateBook)
	routes.Delete("/api/v1/book/:id", middlewares.IsAuthorized(), api.DeleteBook)
}
