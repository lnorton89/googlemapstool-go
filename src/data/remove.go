package data

import (
	_ "github.com/mattn/go-sqlite3"
)

// ParseMultiLineEntryAndInsert parses a fyne.NewMultiLineEntry widget into a slice of strings for each line
// and inserts each line as a new entry in the "destinations" table of the SQLite database.
func RemoveEntry(dbPath string) error {
	//remove entry

	return nil
}
