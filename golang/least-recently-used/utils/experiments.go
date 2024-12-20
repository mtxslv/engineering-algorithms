package utils

import (
	"fmt"
)

func SimpleSimulation(){
	cacheSize := 3
	lru := NewLRUCacheV0(cacheSize)
	lruMisses := 0

	requests := []int{4, 4, 1, 2, 4, 2, 3, 4, 5, 2, 2, 1, 2, 2}
	
	// First Epoch
	for _, el := range []int{1, 2, 1, 5}{
		lru.Put(el, 10*el)
	}
	for _, req := range requests {
		if !lru.Get(req) {
			lruMisses++
			lru.Put(req,10*req)
		}
	}
	fmt.Printf("Cache Size: %d, Requests: %d, LRU Misses: %d\n", cacheSize, len(requests), lruMisses)
}

func SimulateRequests(cacheSize int, requests []int) {
	lru := NewLRUCacheV0(cacheSize)
	lruMisses := 0

	for _, req := range requests {
		if !lru.Get(req) {
			lruMisses++
			lru.Put(req, req)
		}
	}

	fmt.Printf("Cache Size: %d, Requests: %d, LRU Misses: %d\n", cacheSize, len(requests), lruMisses)
}

func SimulateRequestsVaryingInputSize(cacheSize int, inputSizes []int) {
	for _, size := range inputSizes {
		requests := make([]int, size)
		for i := 0; i < size; i++ {
			requests[i] = i % (cacheSize + 1)
		}

		lru := NewLRUCacheV0(cacheSize)
		lruMisses := 0

		for _, req := range requests {
			if !lru.Get(req) {
				lruMisses++
				lru.Put(req, req)
			}
		}

		fmt.Printf("Cache Size: %d, Input Size: %d, LRU Misses: %d\n", cacheSize, size, lruMisses)
	}
}

func SimulationWithRandomness(){

	cacheSizes := []int{3,6,9,12}
	var multiplier int
	var maxLim int
	var numberOfRequests int
	var requests []int

	for _, aCacheSize := range cacheSizes {
		lru := NewLRUCacheV0(aCacheSize)
		multiplier = 2
		for multiplier <= 4 {
			maxLim = int(1.5*float32(aCacheSize))		
			numberOfRequests = multiplier*aCacheSize
			requests = RandomIntSlice(numberOfRequests,maxLim)
			msg := CheckMisses(lru,requests)
			result := fmt.Sprintf("CACHE: %d | %d REQUESTS | RandLimMax: %d | %s\n",aCacheSize,numberOfRequests,maxLim,msg)
			fmt.Printf("%s",result)
			multiplier++
		}
	}
}

/* wordLen can be 6, listLen can be 10.

 requestsNum is length of requests (N), 
 while capacity is cache size (K)
*/
 func SimulationLRUOPT(wordLen, listLen, requestsNum, capacity int){
	// First of all generate ids [strings]
	words := generateRandomStrings(wordLen,listLen)
	requests := generateRandomSequenceStr(wordLen,words)

	// Now the caches...
	lru := NewLRUCacheV1(capacity)
	opt := NewOPTCache(capacity,requests)

	// ... and their misses
	lruMisses := 0
	optMisses := 0

	// Now let's iterate on the requests
	for _, key := range requests { 
		// First on LRU
		_, okLru := lru.Get(key)
		if !okLru { // not found
			lru.Put(key, 1.6180) 
			lruMisses++
		}
		// Now Opt
		_, okOpt := opt.Get(key)
		if !okOpt { // not found
			opt.Put(key, 2.71828) 
			optMisses++
		}
	}
	fmt.Printf("LRU misses: %d | OPT misses: %d\n", lruMisses, optMisses)
	fmt.Printf("COMPETITIVENESS: %.3f\n", float32(lruMisses)/float32(optMisses))
}