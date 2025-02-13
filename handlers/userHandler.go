package handlers

import (
	"context"

	"github.com/MatTwix/Go-React/database"
	"github.com/MatTwix/Go-React/models"
	"github.com/gofiber/fiber/v3"
)

// Получение всех пользователей
func GetUsers(c fiber.Ctx) error {
	rows, err := database.DB.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка запроса к базе данных"})
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Ошибка обработки данных"})
		}
		users = append(users, user)
	}

	return c.JSON(users)
}

// Создание нового пользователя
func CreateUser(c fiber.Ctx) error {
	user := new(models.User)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	_, err := database.DB.Exec(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка добавления пользователя"})
	}

	return c.JSON(user)
}
