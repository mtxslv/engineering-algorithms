package utils

import (
	"fmt"
)

type HashTableOpenAddressingDoubleHashing struct {
	table [HashTableSize]*HashEntry
}

// Hash function to calculate the index
func hashFuncDoubleHashing(key string, i int) int {
	k := 0
	for _, char := range key {
		k += int(char)
	}
	// h1(k)
	h_1_k := k % HashTableSize // HashTableSize is 'm'
	// h2(k)
	h_2_k := 1 + (k % (HashTableSize - 1))
	// h(k,i)
	h_k_i := (h_1_k + i*h_2_k) % HashTableSize
	return h_k_i
}

// Insert a key-value pair into the hash table
func (h *HashTableOpenAddressingDoubleHashing) HashInsert(key string, value float64) (int, bool) {
	i := 0
	// Linear probing for open addressing
	for i < HashTableSize {
		q := hashFuncDoubleHashing(key,i)
		if h.table[q] == nil {
			h.table[q] = &HashEntry{key: key, value: value}	// Insert new entry
			return q, true // Insert successfully
		} else {
			i++
		}
	}
	return -1, false // hash table overflow
}

// Search for a value by key
func (h *HashTableOpenAddressingDoubleHashing) Search(key string) (int, bool) {
	i := 0
	// Linear probing for open addressing
	for i < HashTableSize {
		q := hashFuncDoubleHashing(key,i)
		if h.table[q].key == key {
			return q, true
		} else {
			i++
		}
	}
	return -1, false // hash table overflow
}

func (h *HashTableOpenAddressingDoubleHashing) Get(key string) (*HashEntry, bool) {
	position, found := h.Search(key)
	if found {
		return h.table[position], found
	} else {
		return nil, found
	}
} 

// Display the contents of the hash table
func (h *HashTableOpenAddressingDoubleHashing) Display() {
	for i, entry := range h.table {
		if entry != nil {
			fmt.Printf("Index %d: %s => %.2f\n", i, entry.key, entry.value)
		} else {
			fmt.Printf("Index %d: nil\n", i)
		}
	}
}