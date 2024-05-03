package data

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"mapscreator/src/utils"
)

//go:embed data.csv
var csvData []byte

type DestinationsData struct {
	Destinations []string
}

func Load() *DestinationsData {
	var records [][]string
	var destinations []string
	reader := csv.NewReader(bytes.NewReader(csvData))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			utils.ShowErrorDialog(string(fmt.Sprintf("%s%v", "Error loading data: }", err)))
			return nil
		}

		records = append(records, record)
	}

	for _, record := range records {
		destinations = append(destinations, record...)
	}

	return &DestinationsData{Destinations: destinations}
}
