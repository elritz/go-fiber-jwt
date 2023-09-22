package main

import (
	"log"

	"github.com/elritz/go-fiber-jwt/database"
	"github.com/elritz/go-fiber-jwt/router"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	println("Hello, World!")
	app := fiber.New()

	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
