package main

import (
	"CardService/config"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// go run ./cmd/migrator
func main() {
	cfg := config.LoadMigrator()

	m, err := migrate.New("file://migrations", cfg.DsnLoad())
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no migrations to apply — database already up to date")
			return
		}
		log.Fatalf("migration failed: %v", err)
	}
}