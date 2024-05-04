package data

import (
	"mapscreator/src/utils"

	"database/sql"
	"fmt"
)

// OpenDatabaseConnection opens a connection to the SQLite database specified by dbPath.
func OpenDatabaseConnection(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "DATABASE ERROR: ", err)))
	}

	if err := db.Ping(); err != nil {
		db.Close()
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "DATABASE ERROR: ", err)))
	}

	return db
}
