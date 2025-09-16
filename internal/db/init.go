package db

import (
	"database/sql" // Provides SQL database functionality
	"fmt"          // For printing messages
	"log"          // For logging errors

	_ "github.com/lib/pq" // Import pq for PostgreSQL driver
)

// Custom type to exclude chance of having priority values beside 3: low, mid, high
func createEnum(db *sql.DB) {
	// Define the SQL query for creating enum for priority field
	createEnumSQL := `
DO $$ BEGIN
    CREATE TYPE priority_type AS ENUM ('low', 'mid', 'high');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;`

	// Execute the SQL query
	_, err := db.Exec(createEnumSQL)
	if err != nil {
		log.Fatalf("Failed to create type: %v", err)
	}
	fmt.Println("Type created successfully.")
}

func createTable(db *sql.DB) {
	// Define the SQL query for creating a new table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		description VARCHAR(127),
		completed BOOLEAN,
		due_date TIMESTAMP,
		priority priority_type
	);`

	// Execute the SQL query
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table created successfully.")
}

func InitDB() *sql.DB {
	// Define the connection string with PostgreSQL credentials
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping to confirm connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgreSQL successfully!")

	createEnum(db)
	createTable(db)

	return db
}

func DBClose(db *sql.DB) {
	db.Close()
}
