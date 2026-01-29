package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func NewPostgres(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Println("Database connection failed:", err)
	}

	return db
}
