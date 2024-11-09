package main

import (
	// "fmt"
	"utils/utils"
)

func main(){
	// Create a new hash table
	hashTable := utils.NewHashTable()

	// Load fruits and prices
	fruitsAndPrices := utils.LoadCsv("./fruits-prices.csv")
	
	// Insert fruits and Prices into a regular Hash Table
	for _, item := range(fruitsAndPrices){
		name, price := item.GetName(), item.GetPrice()
		hashTable.Insert(name, price)
	} 
	
	hashTable.Display()

}