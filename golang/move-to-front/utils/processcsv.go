// Huge thanks to SyntaxRules for the snippet
// https://stackoverflow.com/questions/24999079/reading-csv-file-in-go

package utils

import (
    "encoding/csv"
    "log"
    "os"
	"fmt"
	// "strconv"
)


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

func ProcessCsv(records [][]string) []Item {
	items := make([]Item, len(records))
	fmt.Printf("%d records in total.\n",len(records))
	for i, row := range records {
		author := row [0]
		number := row [1]
		title := row [2]
		length := row [3]
		albumName := row [4]
		fmt.Printf("%d : AUTHOR = %s | NUMBER = %s | TITLE = %s | LENGTH = %s | ALBUM = %s\n\n",i, author, number,title, length, albumName)
		item := Item{author, number,title,length,albumName}
		items[i] = item
	}
	return items
}

func LoadCsv(filePath string) []Item {
    records := ReadCsvFile(filePath)
    items := ProcessCsv(records)
    return items
}

// func (item Item) GetItem() float64 {
//     return item.price
// }

// func (item Item) GetName() string {
//     return item.name
// }