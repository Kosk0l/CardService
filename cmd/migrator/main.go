package main

import (
	//"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// migrate - утилита
// go run ./cmd/migrator --dsn="postgres://postgresCrud:qwerty@localhost:5433/postgresCrud?sslmode=disable" --migrations-path=./migrations
func main() {
	
}