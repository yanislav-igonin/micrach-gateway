package controllers

import (
	Config "micrach-gateway/config"
	Models "micrach-gateway/db/models"

	Repositories "micrach-gateway/db/repositories"
	Dto "micrach-gateway/dto"

	"github.com/gofiber/fiber/v2"
)

func AddBoard(c *fiber.Ctx) error {
	headerApiKey := c.Get("Authorization")
	if headerApiKey != Config.App.ApiKey {
		return fiber.NewError(fiber.StatusUnauthorized, "api key must be provided")
	}

	body := new(Dto.CreateBoardRequest)
	err := c.BodyParser(&body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	board := new(Models.Board)
	board.Shortcut = body.ID
	board.Url = body.Url
	board.Name = body.Name
	err = Repositories.Boards.Create(board)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"ok": true})
}
