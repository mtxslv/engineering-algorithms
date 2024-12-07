package main

import (
	"fmt"
	"os"
	"utils/utils"
)

func main() {

	// Check if path to csv file was provided
	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		fmt.Printf("Usage: \n./main <path-to-csv-file>\n")
		return
	} 
	
	// If everything is fine, code follows...

	// Load CSV
	filePath := argsWithProg[1]
    trackList := utils.LoadCsv(filePath)

	// Create a new linked list.
	ll := utils.New()

	// Add items to the linked list
	for it, track := range trackList {
		ll.Add(&track)
		if it == 5 {
			break
		}
	} 

	fmt.Printf("\n\n")
	// Traverse and print the linked list.
	current := ll.Head
	for current != nil {
		fmt.Printf("%+v\n", *current.Content)
		current = current.Next
	}	
	
	ll.MoveToFront(ll.Tail)

	fmt.Printf("\n\n")
	// Traverse and print the linked list.
	current = ll.Head
	for current != nil {
		fmt.Printf("%+v\n", *current.Content)
		current = current.Next
	}		
}