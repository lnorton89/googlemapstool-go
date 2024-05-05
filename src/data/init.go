package data

import (
	"mapscreator/src/utils"

	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB_PATH = utils.AppSettings().DatabasePath

// InitDB initates a SQLite database in the location specified by dbPath.
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
