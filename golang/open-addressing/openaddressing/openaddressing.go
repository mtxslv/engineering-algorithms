package openaddressing

import (
	"fmt"
)

const hashTableSize = 10

type HashTableOpenAddressing struct {
	table [hashTableSize]*HashEntry
}

type HashEntry struct {
	key   string
	value int
}

// Hash function to calculate the index
func hashFunc(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash*31 + int(char)) % hashTableSize
	}
	return hash
}

// Insert a key-value pair into the hash table
func (h *HashTableOpenAddressing) Insert(key string, value int) {
	index := hashFunc(key)

	// Linear probing for open addressing
	for h.table[index] != nil {
		if h.table[index].key == key {
			h.table[index].value = value // Update if key already exists
			return
		}
		index = (index + 1) % hashTableSize // Move to the next slot
	}

	// Insert new entry
	h.table[index] = &HashEntry{key: key, value: value}
}

// Search for a value by key
func (h *HashTableOpenAddressing) Search(key string) (int, bool) {
	index := hashFunc(key)

	// Linear probing to find the key
	for h.table[index] != nil {
		if h.table[index].key == key {
			return h.table[index].value, true
		}
		index = (index + 1) % hashTableSize
	}
	return 0, false // Key not found
}

// Delete a key-value pair
func (h *HashTableOpenAddressing) Delete(key string) bool {
	index := hashFunc(key)

	// Linear probing to find the key
	for h.table[index] != nil {
		if h.table[index].key == key {
			h.table[index] = nil // Remove the entry

			// Rehash subsequent entries
			nextIndex := (index + 1) % hashTableSize
			for h.table[nextIndex] != nil {
				rehashKey, rehashValue := h.table[nextIndex].key, h.table[nextIndex].value
				h.table[nextIndex] = nil
				h.Insert(rehashKey, rehashValue)
				nextIndex = (nextIndex + 1) % hashTableSize
			}
			return true
		}
		index = (index + 1) % hashTableSize
	}
	return false // Key not found
}

// Display the contents of the hash table
func (h *HashTableOpenAddressing) Display() {
	for i, entry := range h.table {
		if entry != nil {
			fmt.Printf("Index %d: %s => %d\n", i, entry.key, entry.value)
		} else {
			fmt.Printf("Index %d: nil\n", i)
		}
	}
}