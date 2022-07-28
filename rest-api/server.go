package main

import (
	"cesargdd/rest-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	routes.Setup(app)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Listen(":8081")
}
