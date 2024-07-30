package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mchu7797/router_api/controller"
	"github.com/mchu7797/router_api/database"
)

func main() {
	app := fiber.New()
	controller.HookHandler(app)

	err := database.Connect("router_api.db")

	if err != nil {
		panic(err)
	}

	log.Fatal(app.Listen(":2024"))
}
