package utils 

import (
	"testing"
)

func TestRandomIntSlice(t *testing.T){
	listSize := 150
	maximumValue := 12
	list := RandomIntSlice(listSize,maximumValue)
	if len(list)!=listSize{
		t.Errorf("List Size wrong.")
	}
	for _, el := range list {
		if el < 1 || el > maximumValue {
			t.Errorf("Values wrong")
		}
	}
}

func TestCheckMisses(t *testing.T){
	lru := NewLRUCache(3)
	requests := []int{1,2,3}
	msg := CheckMisses(lru,requests)
	if msg != "3 misses"{
		t.Errorf("Number of cache misses wrong")
	}
}