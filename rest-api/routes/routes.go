package routes

import (
	"cesargdd/rest-api/controllers"
	"cesargdd/rest-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	blogApp := app.Group("/")
	blogApp.Post("signin", controllers.SignIn)
	blogApp.Post("login", controllers.Login)

	auth := blogApp.Use(middlewares.IsAuth)
	auth.Get("posts", controllers.Posts)
}
