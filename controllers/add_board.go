package controllers

import (
	Config "micrach-gateway/config"

	"github.com/gofiber/fiber/v2"
)

func AddBoard(c *fiber.Ctx) error {
	headerApiKey := c.Get("Authorization")
	if headerApiKey != Config.App.ApiKey {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	// board := Board{}
	// c.BodyParser(&board)
	// boards[board.ID] = board
	return c.SendString("Hello, World 👋!")
}
