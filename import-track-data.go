package main

import (
	"import-track-data/db"
	"import-track-data/util"
	"os"
)

func main() {
	var file *os.File = util.OpenFile("data.csv")
	var rows = util.ReadCSV(file)
	defer file.Close()

	db.WriteDatabase(rows)
}
