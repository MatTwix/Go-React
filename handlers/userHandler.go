package handlers

import (
	"context"

	"github.com/MatTwix/Go-React/database"
	"github.com/MatTwix/Go-React/models"
	"github.com/gofiber/fiber/v3"
)

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

func GetUser(c fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	err := database.DB.
		QueryRow(context.Background(), "SELECT id, name, email FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Пользователь не найден"})
	}

	return c.JSON(user)
}

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

func UpdateUser(c fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)

	if err := c.Bind().Body(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	_, err := database.DB.
		Exec(context.Background(), "UPDATE users SET name = $1, email = $2, WHERE id = $3", user.Name, user.Email, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка обновления пользователя"})
	}

	return c.JSON(fiber.Map{"message": "Пользователь обновлен"})
}

func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DB.
		Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка удаления пользователя"})
	}

	return c.JSON(fiber.Map{"message": "Пользователь удален"})
}
