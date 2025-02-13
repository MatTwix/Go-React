package main

import (
	"log"
	"os"

	"github.com/MatTwix/Go-React/database"
	"github.com/MatTwix/Go-React/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	PORT := os.Getenv("PORT")

	database.ConnectDB()
	defer database.DB.Close()

	routes.SetupRoutes(app)

	app.Listen(":" + PORT)
}
