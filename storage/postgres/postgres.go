package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER = "postgres"
	DB_PASS = "Enrique1!"
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_NAME = "api_prueba"
	DB_SSL  = "disable"
)

var (
	db *sql.DB
)

func NewPostgresDB() (*sql.DB, error) {
	userName := DB_USER
	password := DB_PASS
	host := DB_HOST
	port := DB_PORT
	database := DB_NAME
	sslMode := DB_SSL

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		userName,
		password,
		host,
		port,
		database,
		sslMode,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
