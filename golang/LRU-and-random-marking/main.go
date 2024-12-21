package main

import (
	"fmt"
	"os"
	"utils/utils"
)

func main(){

	// Check if parameters were provided
	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		fmt.Printf("Usage: \n./main <path-to-txt-file>\n")
		return
	} 

	filePath := argsWithProg[1]
	requests := utils.LoadLines(filePath)
	
	// Now the caches...
	capacity := 15 // universe of 39 items requested
	lru := utils.NewLRUCacheV1(capacity)
	rmc, err := utils.NewRandomMarkingCache(capacity)
	if err != nil {
		panic(err)
	}

	// ... and their misses
	lruMisses := 0
	rmcMisses := 0

	// Now let's iterate on the requests
	for _, key := range requests { 
		// First on LRU
		_, okLru := lru.Get(key)
		if !okLru { // not found
			lru.Put(key, 1.6180) 
			lruMisses++
		}
		// Now Opt
		_, okOpt := rmc.Get(key)
		if !okOpt { // not found
			rmc.Put(key, 2.71828) 
			rmcMisses++
		}
	}

	// Display results
	fmt.Printf("Total Requests (N): %d | Cache Size (K): %d | ", len(requests), capacity)
	fmt.Printf("LRU Misses: %d | RMC Misses: %d | ", lruMisses, rmcMisses)
	fmt.Printf("Competitiveness (LRU/OPT): %.3f\n", float32(lruMisses)/float32(rmcMisses))

}