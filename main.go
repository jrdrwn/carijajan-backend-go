package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	InitDatabase()

	app := fiber.New()
	PORT := os.Getenv("PORT")
	api := app.Group("/api")

	InitRoutes(api)

	app.Listen(":" + PORT)
}
