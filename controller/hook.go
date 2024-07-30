package controller

import "github.com/gofiber/fiber/v2"

func HookHandler(app *fiber.App) {
	app.Get("/login", Login)
	app.Get("/intranet", GetIntranet)
	app.Get("/network", GetNetwork)
	app.Get("/user", GetUser)

	app.Get("/", Index)
}
