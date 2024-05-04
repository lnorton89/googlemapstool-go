package routes

import (
	"fmt"
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

	bindingList.OnSelected = func(selected int) {
		destinations, err := viewModel.Destinations.Get()
		if err != nil {
			fmt.Println(err)
			return
		}
		if selected >= 0 && selected < len(destinations) {
			itemText := destinations[selected]
			fmt.Println("Selected item:", itemText)
		}
	}

	return bindingList
}
