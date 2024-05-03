package components

import (
	"mapscreator/src/routes"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func Tabs() *container.AppTabs {
	return container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), container.NewVBox(
			routes.Home(),
		)),
		container.NewTabItemWithIcon("Destinations", theme.HistoryIcon(), container.NewStack(
			routes.Destinations(),
		)),
	)
}
