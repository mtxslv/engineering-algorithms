package utils


import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func generateTrackList(listSize int) []Item {
	stringSize := 50
	it := 0
	tracklist := make([]Item, listSize)
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
		track := Item{
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


func getSampleSongNames(tracklist []Item, sampleSize int) []string {
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

func offlineOptimize(tracklist *LinkedList, names []string) int {

	type titleCount struct {
		title string
		number int
	}
	
	var requests []titleCount

	// Define how many times each name
	// appear in offline request batch
	counter := make(map[string]int)
	for _, name := range names {
		counter[name]++
	}

	for k := range counter {
		requests = append(requests, titleCount{title:k,number:counter[k]})
	}

	sort.Slice(requests, func(i, j int) bool { return requests[i].number < requests[j].number })
	// fmt.Printf("\n%+v\n",requests)

	// Now build a new list 
	totalCost := 0
	for _, titleCountObj := range requests{
		_, cost := tracklist.SearchAndMoveToFrontWithCostIncurred(titleCountObj.title)
		totalCost += int(cost)
	}

	return totalCost
}

func ExperimentRequests(listSize, requestsNumber int) {
	// Generate tracklist
	trackList := generateTrackList(listSize)

	// Create Linked Lists
	llMTF := New()
	llForesee := New()

	// Add items to them
	for _, track := range trackList {
		// fmt.Print("%+v",track)
		llMTF.Add(&track)
		llForesee.Add(&track)
		// fmt.Printf("%s\n",track.Title)
	} 	
	
	// Generate random requests
	requests := getSampleSongNames(trackList, requestsNumber)

	offlineOptimizationCost := offlineOptimize(llForesee, requests)

	totalCostForesee := offlineOptimizationCost
	totalCostMTF := 0

	for _, name := range requests {
		foundF, costForesee := llForesee.SearchWithCostIncurred(name) 
		foundM, costMTF := llMTF.SearchAndMoveToFrontWithCostIncurred(name) 
		if foundF == nil || foundM == nil {
			fmt.Printf("NOT FOUND. ERROR")
		}
		totalCostForesee += int(costForesee)
		totalCostMTF += int(costMTF)
	}

	fmt.Printf("LIST SIZE: %d . REQUESTS SIZE: %d .RATIO: %.4f\n", listSize, requestsNumber, float32(totalCostMTF)/float32(totalCostForesee))

	// fmt.Printf(" ====== OFFLINE TRACKLIST ====== \n")
	// current := llForesee.Head 
	// for current != nil {
	// 	fmt.Printf("%s\n", current.Content.Title)
	// 	current = current.Next
	// }
	// fmt.Printf(" ======================== \n")
	
	// fmt.Printf(" ====== REQUESTS ====== \n")
	// for _, req := range requests {
	// 	fmt.Printf("%s\n", req)
	// }
	// fmt.Printf(" ======================== \n")

}

func ExperimentListSizeEqualToRequestsWorstCase(listSize, requestsNumber int){

	// Generate tracklist
	trackList := generateTrackList(listSize)

	// Create Linked Lists
	llMTF := New()
	llForesee := New()

	// Add items to them
	for _, track := range trackList {
		// fmt.Print("%+v",track)
		llMTF.Add(&track)
		llForesee.Add(&track)
		// fmt.Printf("%s\n",track.Title)
	} 	

	// Generate string list to keep requests
	requests := make([]string, llMTF.Len)
	// Fill requests 
	it := 0
	current := llMTF.Tail
	for current != nil {
		requests[it] = current.Content.Title
		it++
		current = current.Previous
	}

	// fmt.Printf("%+v",requests)
	totalCostForesee := 0
	totalCostMTF := 0

	for _, name := range requests {
		foundF, costForesee := llForesee.SearchWithCostIncurred(name) // (50*(51))/2
		foundM, costMTF := llMTF.SearchAndMoveToFrontWithCostIncurred(name) // (126+125)*50
		if foundF == nil || foundM == nil {
			fmt.Printf("NOT FOUND. ERROR")
		}
		totalCostForesee += int(costForesee)
		totalCostMTF += int(costMTF)
	}

	fmt.Printf("LIST SIZE: %d . REQUESTS SIZE: %d .RATIO: %.4f\n", listSize, requestsNumber, float32(totalCostMTF)/float32(totalCostForesee))

}