package main

import (
	"log"

	"github.com/Galbeyte1/go-web-dev/database"
	"github.com/Galbeyte1/go-web-dev/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

    log.Fatal(app.Listen(":50051"))
}

