package utils

import (
	"strings"
)

func FormatDestinationsInput(text string) string {
	text = strings.ReplaceAll(strings.TrimSpace(text), " ", "+")
	text = strings.ReplaceAll(strings.TrimSpace(text), ",", "")
	text = strings.ReplaceAll(strings.TrimSpace(text), "\n", "/")

	return text
}
