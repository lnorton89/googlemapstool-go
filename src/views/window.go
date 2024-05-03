package views

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
	app fyne.App
	win fyne.Window
}

func NewWindow() *Window {
	app := app.New()
	win := app.NewWindow("Google Maps Route Tool")
	win.Resize(fyne.NewSize(909, 600))

	return &Window{
		app: app,
		win: win,
	}
}

func (w *Window) ShowAndRun() {
	appData := data.Load()
	if appData == nil {
		utils.ShowErrorDialog("Failed to load data from CSV")
		return
	}

	description := container.NewVBox(
		widget.NewLabel("While Google Maps only allows up to 10 destinations, this tool lets you generate routes with more!"),
		widget.NewLabel("Simply enter your desired destinations below, each on a separate line."),
		widget.NewLabel("Click the button, and we'll create a Google Maps link with your optimized route."),
	)

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter your destinations by name or address (one per line):")
	input.SetMinRowsVisible(25)

	submit := widget.NewButton("Generate Link", func() {
		browser.Open(utils.ReplaceSpace(input.Text))
	})
	submit.Importance = widget.SuccessImportance

	list := List(appData.Destinations)

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), container.NewVBox(
			description,
			input,
			submit,
		)),
		container.NewTabItemWithIcon("Destinations", theme.HistoryIcon(), container.NewStack(
			list,
		)),
	)

	w.win.SetContent(tabs)
	w.win.ShowAndRun()
}
