package controller

import "github.com/gofiber/fiber/v2"

func GetIntranet(c *fiber.Ctx) error {
	return c.SendString("Intranet Handler.")
}
