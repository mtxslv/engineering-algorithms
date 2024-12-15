package utils

import (
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