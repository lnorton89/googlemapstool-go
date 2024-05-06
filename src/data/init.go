package data

import (
	"mapscreator/src/utils"

	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// CreateDB creates a SQLite database.
func CreateDB() {
	var err error
	db := OpenDB()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS destinations (
		id INTEGER PRIMARY KEY,
		address TEXT UNIQUE
	)`)
	if err != nil {
		utils.ShowErrorDialog("DATABASE ERROR", err)
	}

	db.Close()
}

// InitDB checks if an SQLite database exists.
// If no SQLite database exists call CreateDB()
func InitDB() {
	f, err := os.Open(DB_PATH)

	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			CreateDB()
		} else {
			fmt.Println("Error opening file:", err)
		}
	}
}
