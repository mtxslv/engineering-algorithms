package main

import (
	// "fmt"
	"fmt"
	"time"
	"utils/utils"
)

func main(){
	// Timing Insertion variables
	var totalInsertionTimeLinearProbe int64 = 0
	var insertionTimeLinearProbe time.Time
	var totalInsertionTimeDoubleHashing int64 = 0
	var insertionTimeDoubleHashing time.Time	

	// Create two hash tables
	doubleHashTable := &utils.HashTableOpenAddressingDoubleHashing{}
	linearProbeTable := &utils.HashTableOpenAddressingLinearProbing{}

	// Load fruits and prices
	fruitsAndPrices := utils.LoadCsv("./fruits-prices.csv")
	
	// Insert fruits and Prices into a Linear Probe Hash Table
	for _, item := range(fruitsAndPrices){
		name, price := item.GetName(), item.GetPrice()

		insertionTimeLinearProbe = time.Now()
		_, worked := linearProbeTable.HashInsert(name, price)
		if worked{
			totalInsertionTimeLinearProbe += time.Since(insertionTimeLinearProbe).Nanoseconds()	
		} else {
			fmt.Printf("%s not inserted\n",name)
		}
	} 
	
	// Insert Fruits and Prices into Open Addressing With Double Hashing
	for _, item := range(fruitsAndPrices) {
		name, price := item.GetName(), item.GetPrice()
		insertionTimeDoubleHashing = time.Now()
		_, worked := doubleHashTable.HashInsert(name, price)
		if worked{
			totalInsertionTimeDoubleHashing += time.Since(insertionTimeDoubleHashing).Nanoseconds()
		} else {
			fmt.Printf("%s not inserted\n",name)
		}
	}
	
	// Report Insertion Time
	fmt.Printf("Took %d nanoseconds to insert with Linear Probe\n", totalInsertionTimeLinearProbe)
	fmt.Printf("Took %d nanoseconds to insert with Double Hasing\n", totalInsertionTimeDoubleHashing)
	
	////////////////////////////////////////////////////////////////////////////////

	// Search Timing variables
	var totalSearchTimeDoubleHashing int64 = 0
	var searchTimeDoubleHashing time.Time
	var totalSearchTimeLinearProbing int64 = 0
	var searchTimeLinearProbing time.Time	

	// Search for fruits on regular Linear Probe
	for _, item := range(fruitsAndPrices){
		name := item.GetName()

		searchTimeLinearProbing = time.Now()
		_, found := linearProbeTable.Get(name) 
		if found {
			totalSearchTimeLinearProbing += time.Since(searchTimeLinearProbing).Nanoseconds()
		}
	} 	
		
		// Search for fruits on Double Hashing
	for _, item := range(fruitsAndPrices){
		name := item.GetName()
			
		searchTimeDoubleHashing = time.Now()
		_, found := doubleHashTable.Get(name) 
		if found {
			totalSearchTimeDoubleHashing += time.Since(searchTimeDoubleHashing).Nanoseconds()
		}
	} 		

	// Report Search Time
	fmt.Printf("\nTook %d nanoseconds to search with Double Hashing\n", totalSearchTimeDoubleHashing)
	fmt.Printf("Took %d nanoseconds to search with Linear Probing\n", totalSearchTimeLinearProbing)
		
	// Display Hash Table
	// linearProbeTable.Display()

	fmt.Printf("\n\n")
	// Display Open Addressing 
	// doubleHashTable.Display()

}