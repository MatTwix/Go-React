package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error laoding .env file:", err)
	}
	DB_URL := os.Getenv("DB_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, DB_URL)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	DB = pool
	log.Println("Успешное подключение к PostgreSQL")
}
