package util

import (
	"encoding/csv"
	"import-track-data/notification"
	"os"
)

func OpenFile(path string) *os.File {
	file, openFileError := os.Open(path)

	notification.Error(openFileError)

	return file
}

func ReadCSV(csvFile *os.File) [][]string {
	csvReader := csv.NewReader(csvFile)
	rows, parseFileError := csvReader.ReadAll()

	notification.Error(parseFileError)

	return rows
}
