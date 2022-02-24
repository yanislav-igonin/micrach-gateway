package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorsHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	message := fiber.ErrInternalServerError.Error()

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Error()
	} else {
		log.Println("error: ", err)
	}

	return ctx.Status(code).JSON(fiber.Map{
		"error": fiber.Map{"message": message},
	})
}
