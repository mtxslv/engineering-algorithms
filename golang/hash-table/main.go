package main

import (
	"fmt"
	"hashtable/hashtable"
)

func main() {
	// Create a new hash table
	hashTable := hashtable.NewHashTable()

	// Insert some key-value pairs
	hashTable.Insert("apple", 10)
	hashTable.Insert("banana", 20)
	hashTable.Insert("orange", 30)
	hashTable.Insert("grape", 40)

	// Display the hash table
	fmt.Println("Hash Table after insertions:")
	hashTable.Display()

	// Retrieve values
	val, found := hashTable.Get("apple")
	if found {
		fmt.Println("Value for 'apple':", val)
	} else {
		fmt.Println("Key 'apple' not found.")
	}

	// Update value
	hashTable.Insert("banana", 25)
	fmt.Println("Hash Table after updating 'banana':")
	hashTable.Display()

	// Delete a key
	deleted := hashTable.Delete("orange")
	if deleted {
		fmt.Println("Deleted 'orange'")
	} else {
		fmt.Println("'orange' not found.")
	}

	// Display the hash table after deletion
	fmt.Println("Hash Table after deletion:")
	hashTable.Display()
}
