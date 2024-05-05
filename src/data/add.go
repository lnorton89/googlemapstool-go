package data

import (
	"mapscreator/src/utils"

	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

// ParseMultiLineEntryAndInsert parses a fyne.NewMultiLineEntry widget into a slice of strings for each line
// and inserts each line as a new entry in the "destinations" table of the SQLite database.
func ParseMultiLineEntryAndInsert(entry *widget.Entry) error {
	var lines = strings.Split(entry.Text, "\n")
	db := OpenDB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO destinations(address) VALUES(?)")
	if err != nil {
		utils.ShowErrorDialog("DATABASE ERROR", err)

		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	for _, line := range lines {
		if _, err := stmt.Exec(line); err != nil {
			utils.ShowErrorDialog("DATABASE ERROR", err)

			return fmt.Errorf("error executing statement: %v", err)
		}
	}
	db.Close()

	return nil
}
