package data

import (
	"fmt"
	"mapscreator/src/utils"
	"strings"

	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

// ParseMultiLineEntryAndInsert parses a fyne.NewMultiLineEntry widget into a slice of strings for each line
// and inserts each line as a new entry in the "destinations" table of the SQLite database.
func ParseMultiLineEntryAndInsert(entry *widget.Entry, dbPath string) error {
	db := OpenDatabaseConnection(dbPath)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO destinations(address) VALUES(?)")
	if err != nil {
		utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error pinging database: }", err)))

		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	for _, line := range strings.Split(entry.Text, "\n") {
		if _, err := stmt.Exec(line); err != nil {
			utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error pinging database: }", err)))

			return fmt.Errorf("error executing statement: %v", err)
		}
	}
	db.Close()

	return nil
}
