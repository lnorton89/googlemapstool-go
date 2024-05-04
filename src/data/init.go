package data

import (
	"mapscreator/src/utils"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initates a SQLite database in the location specified by dbPath.
func InitDB(dbPath string) {
	var err error
	db := OpenDatabaseConnection(dbPath)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS destinations (
		id INTEGER PRIMARY KEY,
		address TEXT UNIQUE
	)`)
	if err != nil {
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "DATABASE ERROR: ", err)))
	}

	db.Close()
}
