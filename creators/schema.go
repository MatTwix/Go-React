package creators

import "github.com/jackc/pgx/v5/pgxpool"

func CreateTables(DB *pgxpool.Pool) {
	CreateUsersTable(DB)
}
