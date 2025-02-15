package creators

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUsersTable(DB *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := DB.Begin(ctx)
	if err != nil {
		log.Fatal("Ошибка начала транзакции:", err)
	}
	defer tx.Rollback(ctx) // Откат в случае ошибки

	// Проверяем, существует ли таблица users
	var tableExists bool
	err = tx.QueryRow(ctx,
		"SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users');").
		Scan(&tableExists)
	if err != nil {
		log.Fatal("Ошибка проверки существования таблицы users:", err)
	}

	// Если таблицы нет, создаём её
	if !tableExists {
		_, err = tx.Exec(ctx, `
			CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				name TEXT NOT NULL,
				email TEXT UNIQUE NOT NULL
			);
		`)
		if err != nil {
			log.Fatal("Ошибка создания таблицы users:", err)
		}
		log.Println("✅ Таблица users успешно создана")
	}

	// Фиксация транзакции
	err = tx.Commit(ctx)
	if err != nil {
		log.Fatal("Ошибка фиксации транзакции:", err)
	}
}
