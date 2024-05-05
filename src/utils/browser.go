package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenBrowser(addresses string) {
	url := "https://www.google.com/maps/dir/" + addresses
	var err error
	switch runtime.GOOS {
	case "linux", "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	default:
		err = fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	if err != nil {
		fmt.Println("Error opening browser:", err)
	}
}
