// Huge thanks to SyntaxRules for the snippet
// https://stackoverflow.com/questions/24999079/reading-csv-file-in-go

package readcsv

import (
    "encoding/csv"
    "log"
    "os"
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