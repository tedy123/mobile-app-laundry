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

	var email string
	var tanggal string
	err = db.QueryRow("SELECT email, tanggal::text FROM transactions LIMIT 1").Scan(&email, &tanggal)
	if err != nil {
		fmt.Printf("Error or No Data: %v\n", err)
	} else {
		fmt.Printf("Sample Data - Email: [%s], Tanggal Text: [%s]\n", email, tanggal)
	}

    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM transactions").Scan(&count)
    fmt.Printf("Total transactions: %d\n", count)
}
