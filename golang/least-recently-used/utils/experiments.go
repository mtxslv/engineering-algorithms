package utils

import "fmt"

func SimpleSimulation(){
	cacheSize := 3
	lru := NewLRUCache(cacheSize)
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
	lru := NewLRUCache(cacheSize)
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

		lru := NewLRUCache(cacheSize)
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