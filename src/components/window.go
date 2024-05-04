package components

import (
	"mapscreator/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Window struct {
	app fyne.App
	win fyne.Window
}

func NewWindow(title string, width float32, height float32) *Window {
	app := app.New()
	win := app.NewWindow(title)
	win.Resize(fyne.NewSize(width, height))
	utils.AppInstance(win)

	return &Window{
		app: app,
		win: win,
	}
}

func (w *Window) ShowAndRun() {
	w.win.SetContent(Router())
	w.win.ShowAndRun()
}
