package utils

import (
	"strings"
)

func FormatDestinationsInput(text string) string {
	return strings.ReplaceAll(strings.TrimSpace(text), "\n", "/")
}
