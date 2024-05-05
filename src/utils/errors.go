package utils

import "fyne.io/fyne/v2/dialog"

func ShowErrorDialog(message string) {
	dialog.NewInformation("Error", message, AppSettings().App).Show()
}
