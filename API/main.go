package main

import (
	"log"
	"petani_edukasi/database"
	"petani_edukasi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err := app.Listen("192.168.141.215:8080")
	if err != nil {
		log.Fatal(err)
	}
}
