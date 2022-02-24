package controllers

import (
	"log"
	Repositories "micrach-gateway/db/repositories"

	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	boards, err := Repositories.Boards.GetAll()
	log.Println(boards)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return ctx.Render("index", fiber.Map{
		"Title":  "Hello, World!",
		"Boards": boards,
	}, "layout/main")
}
