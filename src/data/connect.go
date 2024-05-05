package data

import (
	"mapscreator/src/utils"

	"database/sql"
)

// OpenDB opens a connection to the SQLite database specified by dbPath.
func OpenDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		utils.ShowErrorDialog("DATABASE ERROR", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		utils.ShowErrorDialog("DATABASE ERROR", err)
	}

	return db
}
