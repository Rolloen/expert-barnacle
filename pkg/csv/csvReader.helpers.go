package csvService

import (
	"encoding/csv"
	"os"
)

const (
	localFilePath = "journaux.csv"
)

const (
	CSV_SOURCE_FILE = "localFile"
)

// ReadCSVFromSources : get the datas of a CSV from a given source
// NOTES: need some minor tweaks depending on how we plan to implement other way to get the csv file
func ReadCSVFromSources(source string) ([][]string, error) {
	var csvData [][]string
	var err error
	switch source {
	case CSV_SOURCE_FILE:
		csvData, err = readCSVFromFile(localFilePath)
	}
	// implement other sources

	return csvData, err

}

func readCSVFromFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return data, nil
}
