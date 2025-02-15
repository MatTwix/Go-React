package main

import (
	"log"
	"os"

	"github.com/MatTwix/Go-React/database"
	"github.com/MatTwix/Go-React/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, // Массив строк
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	routes.SetupRoutes(app)

	err := app.Listen(":" + PORT)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
