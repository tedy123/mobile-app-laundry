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

	rows, err := db.Query(`
		SELECT column_name, data_type 
		FROM information_schema.columns 
		WHERE table_name = 'transactions'
		ORDER BY ordinal_position
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Columns in 'transactions' table:")
	for rows.Next() {
		var name, dtype string
		if err := rows.Scan(&name, &dtype); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("- %s (%s)\n", name, dtype)
	}

	rows2, err := db.Query("SELECT * FROM transactions LIMIT 1")
    if err == nil {
        cols, _ := rows2.Columns()
        if rows2.Next() {
            rawValues := make([]interface{}, len(cols))
            valuePtrs := make([]interface{}, len(cols))
            for i := range rawValues {
                valuePtrs[i] = &rawValues[i]
            }
            rows2.Scan(valuePtrs...)
            for i, col := range cols {
                fmt.Printf("Row Data: %s = %v\n", col, rawValues[i])
            }
        }
        rows2.Close()
    }
}
