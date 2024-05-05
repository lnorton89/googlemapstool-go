package data

import (
	"fmt"
	"mapscreator/src/utils"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initates a SQLite database in the location specified by dbPath.
func CreateDB(dbPath string) {
	var err error
	db := OpenDB(dbPath)
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
	var dbPath = utils.AppSettings().DatabasePath
	f, err := os.Open(dbPath)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			CreateDB(dbPath)
		} else {
			fmt.Println("Error opening file:", err)
		}
	}
}
