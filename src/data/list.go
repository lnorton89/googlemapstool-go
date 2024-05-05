package data

import (
	"mapscreator/src/utils"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// ListDestinations retrieves all entries from the "destinations" table of the SQLite database.
func ListDestinations(db *sql.DB) []string {
	var address string
	var addresses []string

	rows, err := db.Query("SELECT address FROM destinations")
	if err != nil {
		utils.ShowErrorDialog("DATABASE ERROR", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&address)
		if err != nil {
			utils.ShowErrorDialog("DATABASE ERROR", err)
		}
		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		utils.ShowErrorDialog("DATABASE ERROR", err)
	}

	return addresses
}
