package data

import (
	"mapscreator/src/utils"

	"database/sql"
)

var DB_PATH = utils.AppSettings().DatabasePath

// OpenDB opens a connection to the SQLite database.
func OpenDB() *sql.DB {
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		utils.ShowErrorDialog("DATABASE ERROR", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		utils.ShowErrorDialog("DATABASE ERROR", err)
	}

	return db
}
