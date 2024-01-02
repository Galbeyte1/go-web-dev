package main

import (
	"log"

	"github.com/Galbeyte1/go-web-dev/database"
	"github.com/Galbeyte1/go-web-dev/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()
	app := fiber.New()

	routes.Setup(app)

    log.Fatal(app.Listen(":50051"))
}

