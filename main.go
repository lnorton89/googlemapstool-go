package main

import (
	"mapscreator/src/components"
	"mapscreator/src/data"
)

func main() {
	data.InitDB()
	components.NewWindow().ShowAndRun()
}
