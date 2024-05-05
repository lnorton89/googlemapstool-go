package main

import (
	"mapscreator/src/components"
	"mapscreator/src/data"
	"mapscreator/src/utils"

	"fmt"
	"os"
)

func main() {
	AppInit()
	components.NewWindow().ShowAndRun()
}

func AppInit() {
	var dbPath = utils.AppSettings().DatabasePath
	f, err := os.Open(dbPath)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			data.InitDB(dbPath)
		} else {
			fmt.Println("Error opening file:", err)
		}
	}
}
