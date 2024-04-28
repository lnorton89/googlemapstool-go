package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type AppData struct {
	Destinations []string
}

func LoadData(filePath string) *AppData {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var records [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV record:", err)
			return nil
		}
		records = append(records, record)
	}

	// Flatten the slice of slices into a single slice
	var destinations []string
	for _, record := range records {
		destinations = append(destinations, record...)
	}

	return &AppData{Destinations: destinations}
}
