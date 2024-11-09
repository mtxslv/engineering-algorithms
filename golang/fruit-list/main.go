package main

import (
	// "fmt"
	"fmt"
	"time"
	"utils/utils"
)

func main(){
	// Timing variables
	var totalInsertionTimeHash int64
	var insertionTimeHash time.Time
	var totalInsertionTimeOpenAddressing int64
	var insertionTimeOpen time.Time	

	// Create a new hash table
	hashTable := utils.NewHashTable()
	openAddressing := &utils.HashTableOpenAddressingDoubleHashing{}

	// Load fruits and prices
	fruitsAndPrices := utils.LoadCsv("./fruits-prices.csv")
	
	// Insert fruits and Prices into a regular Hash Table
	for _, item := range(fruitsAndPrices){
		name, price := item.GetName(), item.GetPrice()

		insertionTimeHash = time.Now()
		hashTable.Insert(name, price)
		totalInsertionTimeHash += time.Since(insertionTimeHash).Nanoseconds()
	} 
	
	
	// Insert Fruits and Prices into Open Addressing With Double Hashing
	for _, item := range(fruitsAndPrices) {
		name, price := item.GetName(), item.GetPrice()
		insertionTimeOpen = time.Now()
		_, worked := openAddressing.HashInsert(name, price)
		if worked{
			totalInsertionTimeOpenAddressing += time.Since(insertionTimeOpen).Nanoseconds()
		} else {
			fmt.Printf("%s not inserted\n",name)
		}
	}
	
	// Report Insertion Time
	fmt.Printf("Took %d nanoseconds to insert in Hash Table\n", totalInsertionTimeHash)
	fmt.Printf("Took %d nanoseconds to insert with Open Addressing\n", totalInsertionTimeOpenAddressing)
	
	
	// Display Hash Table
	// hashTable.Display()

	fmt.Printf("\n\n")
	// Display Open Addressing 
	// openAddressing.Display()

}