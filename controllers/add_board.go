package controllers

import (
	Config "micrach-gateway/config"
	Models "micrach-gateway/db/models"
	Repositories "micrach-gateway/db/repositories"

	"github.com/gofiber/fiber/v2"
)

func AddBoard(c *fiber.Ctx) error {
	headerApiKey := c.Get("Authorization")
	if headerApiKey != Config.App.ApiKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": fiber.Map{
				"message": "api key required",
			},
		})
	}

	var board Models.Board
	err := c.BodyParser(&board)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"message": "error parsing body",
			},
		})
	}

	err = Repositories.Create(&board)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"message": "error creating board",
			},
		})
	}

	return c.JSON(fiber.Map{"ok": true})
}
