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
	"github.com/gofiber/template/pug"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	engine := pug.New("./views", ".pug")
	app := fiber.New(fiber.Config{
		ErrorHandler: Middlewares.ErrorsHandler,
		Views:        engine,
	})
	Config.Init()
	Db.Init()
	defer Db.Disconnect()
	app.Use(recover.New())

	app.Get("/", Controllers.Index)
	app.Post("/api/boards", Controllers.AddBoard)

	log.Println("all systems nominal")
	log.Panicln(app.Listen(":" + strconv.Itoa(Config.App.Port)))
}
