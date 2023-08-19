package main

import (
	"campfire/pkg/config"
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config.LoadEnv(".env")

	uri := os.Getenv("POSTGRES_URI")

	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)

	if err != nil {
		panic(err)
	}

	m.Down()
	m.Up()
	// or m.Step(2) if you want to explicitly set the number of migrations to run
}
