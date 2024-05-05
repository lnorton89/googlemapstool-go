package utils

import (
	"fmt"
	"strings"
)

func ReplaceSpace(text string) string {
	fmt.Println((strings.ReplaceAll(strings.TrimSpace(text), "\n", "+")))
	return strings.ReplaceAll(strings.TrimSpace(text), "\n", "+")
}
