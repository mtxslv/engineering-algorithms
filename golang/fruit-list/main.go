package main

import (
	// "fmt"
	"fmt"
	"time"
	"utils/utils"
)

func main(){
	// Timing Insertion variables
	var totalInsertionTimeHash int64 = 0
	var insertionTimeHash time.Time
	var totalInsertionTimeOpenAddressing int64 = 0
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
			break
		}
	}
	
	// Report Insertion Time
	fmt.Printf("Took %d nanoseconds to insert in Hash Table\n", totalInsertionTimeHash)
	fmt.Printf("Took %d nanoseconds to insert with Open Addressing\n", totalInsertionTimeOpenAddressing)
	
	////////////////////////////////////////////////////////////////////////////////

	// Search Timing variables
	var totalSearchTimeHash int64 = 0
	var searchTimeHash time.Time
	var totalSearchTimeOpenAddressing int64 = 0
	var searchTimeOpen time.Time	

	// Search for fruits on regular Hash Table
	for _, item := range(fruitsAndPrices){
		name := item.GetName()

		searchTimeHash = time.Now()
		_, found := hashTable.Get(name) 
		if found {
			totalSearchTimeHash += time.Since(searchTimeHash).Nanoseconds()
		} else {
			fmt.Print("NOT FOUND")
			break
		}
	} 	

	// Search for fruits on Open Addressing
	for _, item := range(fruitsAndPrices){
		name := item.GetName()

		searchTimeOpen = time.Now()
		_, found := openAddressing.Get(name) 
		if found {
			totalSearchTimeOpenAddressing += time.Since(searchTimeOpen).Nanoseconds()
		} else {
			fmt.Print("NOT FOUND")
			break
		}
	} 		

	// Report Search Time
	fmt.Printf("\nTook %d nanoseconds to search on Hash Table\n", totalSearchTimeHash)
	fmt.Printf("Took %d nanoseconds to search on Open Addressing\n", totalSearchTimeOpenAddressing)
	fmt.Printf("\n\n")
		
	// Display Hash Table
	hashTable.Display()

	// Display Open Addressing 
	openAddressing.Display()

}