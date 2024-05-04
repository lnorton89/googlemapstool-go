package data

import (
	"database/sql"
	"fmt"
	"mapscreator/src/utils"

	_ "github.com/mattn/go-sqlite3"
)

// OpenDatabaseConnection opens a connection to the SQLite database specified by dbPath.
func OpenDatabaseConnection(dbPath string) *sql.DB {
	// Open the SQLite database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error opening database: }", err)))
	}
	// Ping the database to ensure the connection is alive
	if err := db.Ping(); err != nil {
		db.Close()
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error pinging database: }", err)))
	}

	return db
}

// ListDestinations retrieves all entries from the "destinations" table of the SQLite database.
func ListDestinations(db *sql.DB) []string {
	// Query all entries from the "destinations" table
	rows, err := db.Query("SELECT address FROM destinations")
	if err != nil {
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error querying table: }", err)))
	}
	defer rows.Close()

	// Store the retrieved addresses in a slice
	var addresses []string
	for rows.Next() {
		var address string
		err := rows.Scan(&address)
		if err != nil {
			utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error scanning row: }", err)))
		}
		addresses = append(addresses, address)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error iterating rows: }", err)))
	}

	return addresses
}
