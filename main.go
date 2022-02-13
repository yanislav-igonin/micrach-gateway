package main

import (
	"log"
	Config "micrach-gateway/config"
	Controllers "micrach-gateway/controllers"
	Db "micrach-gateway/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()
	Config.Init()
	Db.Init()

	app.Post("/boards", Controllers.AddBoard)

	log.Println("all systems nominal")
	log.Panicln(app.Listen(":" + strconv.Itoa(Config.App.Port)))
}
