package data

import (
	"mapscreator/src/utils"

	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

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
