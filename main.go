package main

import (
	"log"
	Config "micrach-gateway/config"
	Controllers "micrach-gateway/controllers"
	Db "micrach-gateway/db"
	Middlewares "micrach-gateway/middlewares"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: Middlewares.ErrorsHandler,
	})
	Config.Init()
	Db.Init()
	defer Db.Disconnect()
	app.Use(recover.New())

	app.Post("/api/boards", Controllers.AddBoard)

	log.Println("all systems nominal")
	log.Panicln(app.Listen(":" + strconv.Itoa(Config.App.Port)))
}
