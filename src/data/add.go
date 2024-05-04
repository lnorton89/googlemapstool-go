package data

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

// ParseMultiLineEntryAndInsert parses a fyne.NewMultiLineEntry widget into a slice of strings for each line
// and inserts each line as a new entry in the "destinations" table of the SQLite database.
func ParseMultiLineEntryAndInsert(entry *widget.Entry, dbPath string) error {
	if entry.Text == "" {
		return fmt.Errorf("empty")
	}

	// Retrieve lines from the multi-line entry
	lines := strings.Split(entry.Text, "\n")

	// Open the SQLite database
	db := OpenDatabaseConnection(dbPath)
	defer db.Close()

	// Prepare the SQL statement for insertion
	stmt, err := db.Prepare("INSERT INTO destinations(address) VALUES(?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	// Insert each line as a new entry in the "destinations" table
	for _, line := range lines {
		if _, err := stmt.Exec(line); err != nil {
			return fmt.Errorf("error executing statement: %v", err)
		}
	}
	db.Close()

	return nil
}
