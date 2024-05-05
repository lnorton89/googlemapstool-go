package utils

import (
	"fyne.io/fyne/v2"
)

type Settings struct {
	App          fyne.Window
	WindowTitle  string
	Width        float32
	Height       float32
	DatabasePath string
}

func AppSettings() *Settings {
	return &Settings{
		WindowTitle:  "Google Maps Route Tool",
		Width:        1200,
		Height:       700,
		DatabasePath: "data.db",
	}
}

func (s *Settings) SetApp(app fyne.Window) {
	s.App = app
}
