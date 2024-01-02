package routes

import (
	"github.com/Galbeyte1/go-web-dev/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/", controllers.Hello)
	app.Post("/api/register", controllers.Register)
}