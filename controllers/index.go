package controllers

import (
	Repositories "micrach-gateway/db/repositories"

	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	boards, err := Repositories.Boards.GetAll()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return ctx.Render("index", fiber.Map{
		"Title":  "Hello, World!",
		"Boards": boards,
	}, "layout/main")
}
