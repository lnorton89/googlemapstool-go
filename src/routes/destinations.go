package routes

import (
	"mapscreator/src/data"
	"mapscreator/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Destinations() *fyne.Container {
	appData := data.Load()

	if appData == nil {
		utils.ShowErrorDialog("Failed to load data from CSV")
	}

	return container.NewStack(
		widget.NewList(
			func() int {
				return len(appData.Destinations)
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("template")
			},
			func(i widget.ListItemID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(appData.Destinations[i])
			}),
	)
}
