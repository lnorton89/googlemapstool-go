package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func List(destinations []string) *widget.List {
	list := widget.NewList(
		func() int {
			return len(destinations)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(destinations[i])
		})
	return list
}
