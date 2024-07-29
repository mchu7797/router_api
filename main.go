package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mchu7797/router_api/controller"
	"github.com/mchu7797/router_api/database"
)

func main() {
	app := fiber.New()

	app.Get("/login", controller.Login)
	app.Get("/intranet", controller.GetIntranet)
	app.Get("/network", controller.GetNetwork)
	app.Get("/user", controller.GetUser)

	app.Get("/", controller.Index)

	database.Init("example.db")

	log.Fatal(app.Listen(":2024"))
}
