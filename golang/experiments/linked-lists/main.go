// main.go
package main

import (
	"fmt"
	"linked-lists/linkedlist" // Replace with the actual module name if you're using a Go module.
)

type wordCount struct {
    word string 
    count int 
}

func main() {
	// Instantiate a linked list
	list := linkedlist.New[wordCount]()

	// Add three elements to the list
	list.Add(wordCount{word: "one", count: 1})
	list.Add(wordCount{word: "two", count: 2})
	list.Add(wordCount{word: "three", count: 3})

	// Display length after adding elements
	fmt.Println("Length after adding elements:", list.Length()) // Output: 3

	// Remove the middle element (2)
	err := list.Remove(wordCount{word: "two", count: 2})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Removed element 2")
	}

	// Display length after removing an element
	fmt.Println("Length after removing element:", list.Length()) // Output: 2

	// Print the remaining elements to verify structure
	printList(list)
}

// Helper function to print all elements in the linked list
func printList[T comparable](ll *linkedlist.LinkedList[T]) {
	current := ll.Head()
	for current != nil {
		fmt.Print(current.Value(), " ")
		current = current.Next()
	}
	fmt.Println()
}
