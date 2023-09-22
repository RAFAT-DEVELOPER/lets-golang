package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

func main() {
	// Case 1: Open a database connection
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Case 2: Create a table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INT
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'users' created successfully.")

	// Case 3: Insert data into the table
	insertSQL := "INSERT INTO users (name, age) VALUES (?, ?);"
	result, err := db.Exec(insertSQL, "Alice", 30)
	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, _ := result.LastInsertId()
	fmt.Printf("Inserted row with ID %d\n", lastInsertID)

	// Case 4: Query data from the table
	querySQL := "SELECT id, name, age FROM users;"
	rows, err := db.Query(querySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Query results:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// Case 5: Update data in the table
	updateSQL := "UPDATE users SET age = ? WHERE name = ?;"
	_, err = db.Exec(updateSQL, 31, "Alice")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated age for Alice.")

	// Case 6: Error handling
	_, err = db.Exec("INVALID SQL")
	if err != nil {
		fmt.Println("Error executing SQL:", err)
	}

	// Case 7: Removing the database file (optional)
	err = os.Remove("example.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database file 'example.db' removed.")
}
