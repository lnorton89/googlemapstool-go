package main

import (
	"mapscreator/src/browser"
	"mapscreator/src/data"
	"mapscreator/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Window struct {
	Name   string
	Width  float32
	Height float32
}

func main() {
	window := Window{
		Name:   "Google Maps Route Tool",
		Width:  909,
		Height: 600,
	}

	a := app.New()
	w := a.NewWindow(window.Name)
	w.Resize(fyne.NewSize(window.Width, window.Height))

	appData := data.LoadData("data.csv")
	if appData == nil {
		utils.ShowErrorDialog("Failed to load data from CSV")
		return
	}

	hello := widget.NewLabel("While Google Maps only allows up to 10 destinations, this tool lets you generate routes with more!")
	hello2 := widget.NewLabel("Simply enter your desired destinations below, each on a separate line. Click the button, and we'll create a Google Maps link with your route.")

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter your destinations by name or address (one per line):")
	input.SetMinRowsVisible(25)

	submit := widget.NewButton("Generate Link", func() {
		browser.OpenBrowser(utils.ReplaceSpace(input.Text))
	})
	submit.Importance = widget.SuccessImportance

	list := widget.NewList(
		func() int {
			return len(appData.Destinations)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(appData.Destinations[i])
		})

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), container.NewVBox(
			hello,
			hello2,
			input,
			submit,
		)),
		container.NewTabItemWithIcon("Destinations", theme.HistoryIcon(), container.NewVBox(
			list,
		)),
	)

	w.SetContent(tabs)
	w.ShowAndRun()
}
