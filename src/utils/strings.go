package utils

import (
	"strings"
)

func ReplaceSpace(text string) string {
	return strings.ReplaceAll(strings.TrimSpace(text), "\n", "/")
}
