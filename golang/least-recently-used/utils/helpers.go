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

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func generateRandomStrings(wordLen, listLen int) []string {
	var arr []string
	it := 0
	for it < listLen {
		arr = append(arr, RandStringBytes(wordLen))
		it++
	}
	return arr
}

func generateRandomSequenceStr(sequenceLen int, words []string) []string {
	
	arr := make([]string, sequenceLen)

	for it := range arr {
		arr[it] = words[rand.Intn(len(words))]
	}

	return arr
}