package main

import (
    "fmt"
	"readcsv/readcsv"
)

func main() {
    records := readcsv.ReadCsvFile("./fruits-prices.csv")
    fmt.Println(records)
}
