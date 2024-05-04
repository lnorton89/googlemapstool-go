package data

import (
	"mapscreator/src/utils"

	"database/sql"
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
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error creating table: }", err)))
	}

	db.Close()
}

// OpenDatabaseConnection opens a connection to the SQLite database specified by dbPath.
func OpenDatabaseConnection(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error opening database: }", err)))
	}

	if err := db.Ping(); err != nil {
		db.Close()
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error pinging database: }", err)))
	}

	return db
}

// func RefreshData() {
// 	db := OpenDatabaseConnection(utils.AppSettings().DatabasePath)
// 	// Query your database for the value
// 	var value string
// 	err := db.QueryRow("SELECT value FROM destinations").Scan(&value)
// 	if err != nil {
// 		fmt.Println("Error fetching data:", err)
// 		return
// 	}

// 	// Update data and trigger UI update
// 	data.Value = value
// 	boundValue.Trigger()
// }
