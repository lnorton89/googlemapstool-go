package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type Settings struct {
	DatabasePath string
}

func AppSettings() *Settings {
	return &Settings{
		DatabasePath: "data.db",
	}
}

var App fyne.Window

func AppInstance(Instance fyne.Window) {
	App = Instance
}

func ReplaceSpace(text string) string {
	return strings.ReplaceAll(strings.TrimSpace(text), " ", "+")
}

func ShowErrorDialog(message string) {
	dialog.NewInformation("Error", message, App).Show()
}

func OpenBrowser(addresses string) {
	var err error
	url := "https://www.google.com/maps/dir/" + addresses

	switch runtime.GOOS {
	case "linux", "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	default:
		err = fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	if err != nil {
		ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error opening browser: }", err)))
	}
}
