package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mchu7797/router_api/database"
	"github.com/mchu7797/router_api/database/models"
)

func SearchNetworkConfig(c *fiber.Ctx) error {
	name := c.Query("name")

	var configData models.NetworkConfig
	result := database.Connection.Where("Name = ?", name).First(&configData)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(configData)
}

func GetNetworkConfig(c *fiber.Ctx) error {
	id := c.Params("id")

	var configData models.NetworkConfig
	result := database.Connection.Find(&configData, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&configData)
}

func AddNetworkConfig(c *fiber.Ctx) error {
	return nil
}

func UpdateNetworkConfig(c *fiber.Ctx) error {
	return nil
}

func DeleteNetworkConfig(c *fiber.Ctx) error {
	return nil
}
