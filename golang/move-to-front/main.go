package main

import (
	"fmt"
	"os"
	"strconv"
	"utils/utils"
	"math/rand"
	// "sort"
	// "time"
)

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func generateTrackList(listSize int) []utils.Item {
	stringSize := 50
	it := 0
	tracklist := make([]utils.Item, listSize)
	for it < listSize {
		// Generate Author as a String
		author := RandStringBytes(16)
		// Generate No. as a String, from a Number between 1 and 15
		number := fmt.Sprintf("%d", rand.Intn(15)+1)
		// Generate Title as a String 
		title := RandStringBytes(stringSize)
		// Generate Length as a String 
		// 		"<Number between 0 and 6>:<Number betweem 0 and 59>"
		length := fmt.Sprintf("%d:%d",rand.Intn(7),rand.Intn(59))
		// Generate album-name as a String
		albumName := RandStringBytes(16)
		
		// Create a track
		track := utils.Item{
			Author: author,
			Number: number,
			Title: title,
			Length: length,
			AlbumName: albumName,			
		}

		// Add to track
		tracklist[it] = track

		// Goes to next
		it++
	}

	return tracklist
}

func main() {

	// Check if parameters were provided
	argsWithProg := os.Args
	if len(argsWithProg) != 3 {
		fmt.Printf("Usage: \n./main <list-size> <number-of-requests>\n")
		return
	} 
	
	// If everything is fine, code follows...

	// Define parameters
	listSize, errArg1 := strconv.Atoi(argsWithProg[1])
	requestsNumber, errArg2 := strconv.Atoi(argsWithProg[2])
	if errArg1 != nil {
		panic(errArg1)
	}
	if errArg2 != nil {
		panic(errArg2)
	}
	if listSize < 1 || requestsNumber < 1 {
		fmt.Printf("Arguments must be numeric and positive.\n")
		return
	}

	trackList := generateTrackList(listSize)
	fmt.Printf("====== TRACKLIST ======\n")
	for _, track := range trackList {
		fmt.Printf("%+v\n", track)
	}
	fmt.Printf("============\n")

}