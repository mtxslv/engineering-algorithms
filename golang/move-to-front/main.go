package main

import (
	"fmt"
	"os"
	"utils/utils"
	"math/rand"
	"sort"
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

func getOfflineList(tracklist *utils.LinkedList, names []string) *utils.LinkedList {
	offlineList := utils.New()

	type titleCount struct {
		title string
		number int
	}
	
	requests := make([]titleCount, tracklist.Len)

	// Define how many times each name
	// appear in offline request batch
	current := tracklist.Head
	it := 0
	for current != nil {
		count := 0
		for _, name := range names{
			if name == current.Content.Title{
				count++
			}
		}
		requests[it] = titleCount{title:current.Content.Title,number:count}
		current = current.Next
		it++
	}
	sort.Slice(requests, func(i, j int) bool { return requests[i].number > requests[j].number })
	// fmt.Printf("\n%+v\n",requests)

	// Now build a new list 
	for _, titleCountObj := range requests{
		obj := tracklist.Search(titleCountObj.title)
		offlineList.Add(obj)
	}

	return offlineList

}

func randomChoiceExperiment(ll *utils.LinkedList, trackList []utils.Item, sampleSize int) {

	names := getSampleSongNames(trackList, sampleSize)

	offlineList := getOfflineList(ll, names)

	totalCostForesee := 0
	totalCostMTF := 0

	for _, name := range names {
		_, costForesee := offlineList.SearchAndMoveToFrontWithCostIncurred(name)
		_, costMTF := ll.SearchAndMoveToFrontWithCostIncurred(name)

		totalCostForesee += int(costForesee)
		totalCostMTF += int(costMTF)
	}

	fmt.Printf("RATIO RANDOM CHOICE: %.4f\n", float32(totalCostMTF)/float32(totalCostForesee))
}

func worstCaseScenarioExperiment(ll *utils.LinkedList, sampleSize int) {

	names := make([]string, sampleSize)
	it := 0
	current := ll.Tail
	for it < sampleSize {
		names[it] = current.Content.Title
		current = current.Previous
		it++
	}

	// fmt.Printf("%+v\n", names)

	offlineList := getOfflineList(ll, names)

	// current = offlineList.Tail
	// for current != nil {
	// 	fmt.Printf("OFFLINE LIST: %+v\n", current.Content)
	// 	current = current.Previous
	// }

	totalCostForesee := 0
	totalCostMTF := 0

	for _, name := range names {
		_, costForesee := offlineList.SearchAndMoveToFrontWithCostIncurred(name)
		_, costMTF := ll.SearchAndMoveToFrontWithCostIncurred(name)

		totalCostForesee += int(costForesee)
		totalCostMTF += int(costMTF)
	}

	fmt.Printf("RATIO WORST CASE: %.4f\n", float32(totalCostMTF)/float32(totalCostForesee))
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

	// Create a pair of linked list.
	llRandom := utils.New()
	llWorst := utils.New()

	// Add items to them
	for _, track := range trackList {
		// fmt.Print("%+v",track)
		llRandom.Add(&track)
		llWorst.Add(&track)
	} 

	// Requests Batch will have 50 song names to look for 
	var sampleSize int = 50

	randomChoiceExperiment(llRandom, trackList, sampleSize)
	// worstCaseScenarioExperiment(llWorst)
	worstCaseScenarioExperiment(llWorst,sampleSize)
	
}