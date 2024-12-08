package main

import (
	"fmt"
	"os"
	"utils/utils"
	"math/rand"
    "time"
)

func getSampleSongNames(tracklist []utils.Item, sampleSize int) []string {
	names := make([]string, sampleSize)
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := 0
	for i < sampleSize {
		position := randomizer.Intn(len((tracklist)))
		names[i] = tracklist[position].Title
		i++
	}
	return names
}

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

	// Randomly sample 50 song names to look for 
	names := getSampleSongNames(trackList, 50)


	fmt.Printf("\n\n")

	for _, songName := range names {
		songMetadata, cost := ll.SearchAndMoveToFrontWithCostIncurred(songName)
		if songMetadata != nil {
			fmt.Printf("\n\t Took %d operations (ratio = %.3f) to find the music's metadata (title: %s)", cost, float32(cost)/100.0 ,songName)
		} else {
			fmt.Printf("\n\t Couldn't find your music :(\n [Took %d operations]", cost)
			return
		}
	}
		
	fmt.Println();
}