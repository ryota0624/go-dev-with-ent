package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://0.0.0.0:5432/sample?user=postgres&password=passward&sslmode=disable",
	)
	if err != nil {
		log.Fatalf("failed to migration New: %v", err)
	}
	if err := m.Up(); err != nil {
		log.Fatalf("failed to migration Up: %v", err)
	}
}
