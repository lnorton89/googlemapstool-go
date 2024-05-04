package data

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) error {
	// Open or create the SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Create the "destinations" table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS destinations (
		id INTEGER PRIMARY KEY,
		address TEXT UNIQUE
	)`)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	return nil
}
