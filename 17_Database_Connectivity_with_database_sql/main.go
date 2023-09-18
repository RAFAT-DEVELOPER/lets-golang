package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	// Query and manipulate data
	// ...

	fmt.Println("Database operations completed")
}
