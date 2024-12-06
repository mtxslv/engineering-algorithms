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
	for _, track := range trackList {
		ll.Add(&track)
	} 

	fmt.Printf("\n\n")
	// Traverse and print the linked list.
	// current := ll.Head
	// for current != nil {
	// 	fmt.Printf("%+v\n", *current.Content)
	// 	current = current.Next
	// }	

	// Search for  a given song
	title := "I Wanna Be Yours"
	itemRef := ll.Search(title)
	if itemRef != nil {
		item := *itemRef
		fmt.Printf(
			"AUTHOR: %s | TRACK NUMBER: %s | TITLE: %s | LENGTH: %s | ALBUM: %s\n",
			item.Author,
			item.Number,
			item.Title,
			item.Length,
			item.AlbumName,
		)
	} else {
		fmt.Printf("Song not found \n")
	}
}