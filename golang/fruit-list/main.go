package main

import (
	// "fmt"
	"fmt"
	"utils/utils"
)

func main(){
	// Create a new hash table
	hashTable := utils.NewHashTable()
	openAddressing := &utils.HashTableOpenAddressingDoubleHashing{}

	// Load fruits and prices
	fruitsAndPrices := utils.LoadCsv("./fruits-prices.csv")
	
	// Insert fruits and Prices into a regular Hash Table
	for _, item := range(fruitsAndPrices){
		name, price := item.GetName(), item.GetPrice()
		hashTable.Insert(name, price)
	} 
	
	// Display Hash Table
	hashTable.Display()

	fmt.Printf("\n\n")

	// Insert Fruits and Prices into Open Addressing With Double Hashing
	for _, item := range(fruitsAndPrices) {
		name, price := item.GetName(), item.GetPrice()
		openAddressing.HashInsert(name, price)
	}

	openAddressing.Display()

}