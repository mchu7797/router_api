package controller

import "github.com/gofiber/fiber/v2"

func GetNetwork(c *fiber.Ctx) error {
	return c.SendString("Network Handler.")
}
