package routes

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func Router() *container.AppTabs {
	return container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), container.NewVBox(
			Home(),
		)),
		container.NewTabItemWithIcon("Destinations", theme.HistoryIcon(), container.NewStack(
			Destinations(),
		)),
	)
}
