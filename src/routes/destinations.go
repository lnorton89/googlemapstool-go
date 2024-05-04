package routes

import (
	"mapscreator/src/data"
	"mapscreator/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type DestinationsViewModel struct {
	Destinations binding.StringList
}

var viewModel = &DestinationsViewModel{
	Destinations: binding.NewStringList(),
}

func RefreshData() {
	destinations := data.ListDestinations(data.OpenDatabaseConnection(utils.AppSettings().DatabasePath))
	viewModel.Destinations.Set(destinations)
}

func Destinations() fyne.CanvasObject {
	database := data.OpenDatabaseConnection(utils.AppSettings().DatabasePath)
	destinations := data.ListDestinations(database)

	viewModel.Destinations.Set(destinations)

	bindingList := widget.NewListWithData(
		viewModel.Destinations,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		},
	)

	return bindingList
}
