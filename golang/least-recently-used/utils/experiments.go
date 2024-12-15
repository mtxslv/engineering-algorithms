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