package routes

import (
	"mapscreator/src/data"
	"mapscreator/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Destinations() *fyne.Container {
	database := data.OpenDatabaseConnection(utils.AppSettings().DatabasePath)
	destinations := data.ListDestinations(database)

	return container.NewStack(
		widget.NewList(
			func() int {
				return len(destinations)
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("template")
			},
			func(i widget.ListItemID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(destinations[i])
			}),
	)
}
