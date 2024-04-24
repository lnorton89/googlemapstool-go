package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Window struct {
	Name   string
	Width  int
	Height int
}

func main() {
	window := Window{
		Name:   "Google Maps Route Tool",
		Width:  909,
		Height: 600,
	}

	a := app.New()
	w := a.NewWindow(window.Name)
	w.Resize(fyne.NewSize(float32(window.Width), float32(window.Height)))
	w.Icon()

	hello := widget.NewLabel("While Google Maps only allows up to 10 destinations, this tool lets you generate routes with more!")
	hello2 := widget.NewLabel("Simply enter your desired destinations below, each on a separate line. Click the button, and we'll create a Google Maps link with your route.")

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter your destinations by name or address (one per line):")
	input.SetMinRowsVisible(25)

	submit := widget.NewButton("Generate Link", func() {
		openBrowser(processString(input.Text))
	})
	submit.Importance = widget.SuccessImportance

	w.SetContent(container.NewVBox(
		hello,
		hello2,
		input,
		submit,
	))

	w.ShowAndRun()
}

func openBrowser(addresses string) {
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

func processString(text string) string {
	lines := strings.Split(text, "\n")
	processed := strings.Join(replaceSpace(lines), "/")

	return processed
}

func replaceSpace(lines []string) []string {
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, " ", "+")
	}

	return lines
}
