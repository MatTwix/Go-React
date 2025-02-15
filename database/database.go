package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/MatTwix/Go-React/creators"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() {
	// Загружаем .env файл
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DB_URL не задан в .env файле")
	}

	// Контекст с таймаутом 5 секунд для подключения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Подключение к БД
	pool, err := pgxpool.New(ctx, DB_URL)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	DB = pool
	log.Println("✅ Успешное подключение к PostgreSQL")

	creators.CreateTables(DB)
}
