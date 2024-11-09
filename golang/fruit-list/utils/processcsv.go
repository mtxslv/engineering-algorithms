// Huge thanks to SyntaxRules for the snippet
// https://stackoverflow.com/questions/24999079/reading-csv-file-in-go

package utils

import (
    "encoding/csv"
    "log"
    "os"
	"strconv"
)

type ItemPrice struct {
	name string
	price float64
}

func ReadCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func ProcessCsv(records [][]string) []ItemPrice {
	items := make([]ItemPrice, len(records))
	for i, row := range records {
		price, _ := strconv.ParseFloat(row[1], 64)
		item := ItemPrice{row[0],price}
		items[i] = item
	}
	return items
}

func LoadCsv(filePath string) []ItemPrice {
    records := ReadCsvFile(filePath)
    items := ProcessCsv(records)
    return items
}