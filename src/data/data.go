package data

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
)

//go:embed data.csv
var csvData []byte

type AppData struct {
	Destinations []string
}

func Load() *AppData {
	reader := csv.NewReader(bytes.NewReader(csvData))
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
