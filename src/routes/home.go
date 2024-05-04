package routes

import (
	"mapscreator/src/data"
	"mapscreator/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Home() *fyne.Container {
	description := container.NewVBox(
		widget.NewLabel("While Google Maps only allows up to 10 destinations, this tool lets you generate routes with more!"),
		widget.NewLabel("Simply enter your desired destinations below, each on a separate line."),
		widget.NewLabel("Click the button, and we'll create a Google Maps link with your optimized route."),
	)

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter your destinations by name or address (one per line):")
	input.SetMinRowsVisible(25)

	submit := widget.NewButton("Generate Link", func() {
		utils.OpenBrowser(utils.ReplaceSpace(input.Text))
		data.ParseMultiLineEntryAndInsert(input, utils.AppSettings().DatabasePath)
	})
	submit.Importance = widget.SuccessImportance

	return container.NewVBox(
		description,
		input,
		submit,
	)
}
