package main

import (
	Config "micrach-gateway/config"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

type Board struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func main() {
	app := fiber.New()
	Config.Init()

	boards := make(map[string]Board)

	app.Post("/boards", func(c *fiber.Ctx) error {
		headerApiKey := c.Get("Authorization")
		if headerApiKey != Config.App.ApiKey {
			return c.Status(401).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		board := Board{}
		c.BodyParser(&board)
		boards[board.ID] = board
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":" + strconv.Itoa(Config.App.Port))
}
