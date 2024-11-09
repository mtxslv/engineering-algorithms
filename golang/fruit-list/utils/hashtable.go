package utils

import (
	"fmt"
)

// Node represents a node in the linked list for chaining
type Node struct {
	key   string
	value float64
	next  *Node
}

// HashTable struct that holds an array of linked lists (chains)
type HashTable struct {
	buckets [HashTableSize]*Node
}

// NewHashTable creates a new hash table
func NewHashTable() *HashTable {
	return &HashTable{}
}

// hashFunction generates an index based on the key
func hashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % HashTableSize
}

// Insert adds a key-value pair to the hash table
func (ht *HashTable) Insert(key string, value float64) {
	index := hashFunction(key)
	newNode := &Node{key: key, value: value, next: nil}

	// If no entry exists at this index, place the new node here
	if ht.buckets[index] == nil {
		ht.buckets[index] = newNode
	} else {
		// Collision occurred, so use chaining to add it to the linked list
		current := ht.buckets[index]
		for current != nil {
			if current.key == key {
				// Update the value if the key already exists
				current.value = value
				return
			}
			if current.next == nil {
				break
			}
			current = current.next
		}
		// Add the new node at the end of the linked list
		current.next = newNode
	}
}

// Get retrieves a value by key from the hash table
func (ht *HashTable) Get(key string) (float64, bool) {
	index := hashFunction(key)
	current := ht.buckets[index]

	// Traverse the linked list at this bucket to find the key
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	// Key not found
	return 0, false
}

// Delete removes a key-value pair from the hash table
func (ht *HashTable) Delete(key string) bool {
	index := hashFunction(key)
	current := ht.buckets[index]
	var prev *Node

	// Traverse the linked list to find and delete the node
	for current != nil {
		if current.key == key {
			if prev == nil {
				// Node is at the head of the list
				ht.buckets[index] = current.next
			} else {
				// Node is in the middle or end of the list
				prev.next = current.next
			}
			return true
		}
		prev = current
		current = current.next
	}
	// Key not found
	return false
}

// Display prints the entire hash table for debugging
func (ht *HashTable) Display() {
	for i, bucket := range ht.buckets {
		fmt.Printf("Bucket %d: ", i)
		current := bucket
		for current != nil {
			fmt.Printf("%s=%.3f -> ", current.key, current.value)
			current = current.next
		}
		fmt.Println("nil")
	}
}