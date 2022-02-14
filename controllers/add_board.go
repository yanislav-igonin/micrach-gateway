package controllers

import (
	Config "micrach-gateway/config"
	Models "micrach-gateway/db/models"

	// Repositories "micrach-gateway/db/repositories"
	Db "micrach-gateway/db"
	Dto "micrach-gateway/dto"
	"time"

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

	var body Dto.CreateBoardRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"message": "error parsing body",
			},
		})
	}

	var board Models.Board
	board.ID = body.ID
	board.Url = body.Url
	board.Description = body.Description
	board.IsUp = true
	board.LastPing = time.Now().UnixMilli()
	err = Db.CreateBoard(&board)
	// err = Repositories.Boards.Create(&board)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"message": "error creating board",
			},
		})
	}

	return c.JSON(fiber.Map{"ok": true})
}
