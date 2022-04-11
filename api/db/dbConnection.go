package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DBConnection() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DatabaseURL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db, err
}
