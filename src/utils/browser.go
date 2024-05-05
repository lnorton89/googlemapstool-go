package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenBrowser(addresses string) {
	var err error
	url := "https://www.google.com/maps/dir/" + FormatDestinationsInput(addresses)

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
