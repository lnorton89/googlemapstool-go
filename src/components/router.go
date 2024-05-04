package components

import (
	"mapscreator/src/routes"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func Router() *container.AppTabs {
	HomeTab := container.NewTabItemWithIcon("Home", theme.HomeIcon(), container.NewVBox(
		routes.Home(),
	))
	DestinationsTab := container.NewTabItemWithIcon("Destinations", theme.HistoryIcon(), container.NewStack(
		routes.Destinations(),
	))

	return container.NewAppTabs(
		HomeTab,
		DestinationsTab,
	)
}
