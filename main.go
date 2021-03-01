package main

import (
	"github.com/barandemirbas/go-jwt-server/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	routes.SetRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
