package utils

import (
	"fmt"

	"fyne.io/fyne/v2/dialog"
)

func ShowErrorDialog(message string, err error) {
	dialog.NewInformation("Error", FormatError(message, err), AppSettings().App).Show()
}

func FormatError(text string, err error) string {
	return string(fmt.Sprintf("%s%v", text+" : ", err))
}
