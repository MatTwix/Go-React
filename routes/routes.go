package routes

import (
	"github.com/MatTwix/Go-React/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", handlers.GetUsers)
	app.Post("/users", handlers.CreateUser)
}
