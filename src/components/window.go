package components

import (
	"mapscreator/src/data"
	"mapscreator/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Window struct {
	app fyne.App
	win fyne.Window
}

func NewWindow() *Window {
	var settings = utils.AppSettings()
	data.InitDB()
	app := app.New()
	win := app.NewWindow(settings.WindowTitle)
	win.Resize(fyne.NewSize(settings.Width, settings.Height))
	settings.SetApp(win)

	return &Window{
		app: app,
		win: win,
	}
}

func (w *Window) ShowAndRun() {
	w.win.SetContent(Router())
	w.win.ShowAndRun()
}
