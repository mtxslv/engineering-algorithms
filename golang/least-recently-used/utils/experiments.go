package utils

import "fmt"

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