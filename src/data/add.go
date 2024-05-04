package data

import (
	"database/sql"
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

// ParseMultiLineEntryAndInsert parses a fyne.NewMultiLineEntry widget into a slice of strings for each line
// and inserts each line as a new entry in the "destinations" table of the SQLite database.
func ParseMultiLineEntryAndInsert(entry *widget.Entry, dbPath string) error {
	// Retrieve lines from the multi-line entry
	lines := strings.Split(entry.Text, "\n")

	// Open the SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
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

	return nil
}
