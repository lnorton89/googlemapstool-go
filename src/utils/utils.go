package utils

import "strings"

func ReplaceSpace(text string) string {
	return strings.ReplaceAll(strings.TrimSpace(text), " ", "+")
}

func ShowErrorDialog(message string) {
	// Show error dialog to the user
}
