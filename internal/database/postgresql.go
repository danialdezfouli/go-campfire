package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetPostgres() *sql.DB {
	return db
}

func CreatePostgresConnection() error {

	uri := os.Getenv("POSTGRES_URI")

	if uri == "" {
		log.Fatal("You must set your 'POSTGRES_URI' environmental variable.")
	}

	var err error
	db, err = sql.Open("postgres", uri)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

func ClosePostgresConnection() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
