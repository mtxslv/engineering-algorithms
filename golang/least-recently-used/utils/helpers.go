package utils

import (
	"fmt"
	"math/rand"
)

// Generate a list of size listSize where each term
// is a random integer between 1 and maximumValue.
func RandomIntSlice(listSize, maximumValue int) []int {
	list := make([]int, listSize)
	it := 0
	for it < listSize{
		list[it] = 1 + rand.Intn(maximumValue)
		it++
	}
	return list
}

func CheckMisses(lru *LRUCacheV0, requests []int) string {
	lruMisses := 0

	for _, req := range requests {
		if !lru.Get(req) {
			lruMisses++
			lru.Put(req,10*req)
		}
	}
	msg := fmt.Sprintf("%d misses", lruMisses)
	return msg
}