package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=laundry port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM transactions LIMIT 0")
	if err != nil {
		log.Fatal(err)
	}
	cols, _ := rows.Columns()
	fmt.Printf("Columns: %v\n", cols)
}
