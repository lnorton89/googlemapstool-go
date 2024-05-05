package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenBrowser(addresses string) {
	var err error
	url := "https://www.google.com/maps/dir/" + addresses

	switch runtime.GOOS {
	case "linux", "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/C", "start", url).Start()
	default:
		ShowErrorDialog(string(fmt.Sprintf("%s%v", "APP ERROR: unsupported platform - ", runtime.GOOS)))
	}
	if err != nil {
		ShowErrorDialog(string(fmt.Sprintf("%s%v", "APP ERROR: ", err)))
	}
}
