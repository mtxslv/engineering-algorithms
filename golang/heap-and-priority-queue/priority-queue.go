package main

import (
	"fmt"
	"math"
	"errors"
)

type heap struct {
	heap []int
	heap_size int
}

func parent(i int) int  { return (i - 1) / 2 }
func left(i int) int    { return 2*i + 1 }
func right(i int) int   { return 2*i + 2 }

// Max heapify function
func max_heapify(A *heap, i int) {
    var l = left(i)
    var r = right(i)
    var largest = i // Initially, assume 'i' is the largest

    // Ensure 'l' is within bounds and check if left child is larger
	if l < A.heap_size && A.heap[l] > A.heap[i] {
		largest = l
	}
	if r < A.heap_size && A.heap[r] > A.heap[largest] {
		largest = r
	}

    // If 'largest' is not 'i', swap and continue heapifying
	if largest != i {
		aux := A.heap[i]
		A.heap[i] = A.heap[largest]
		A.heap[largest] = aux
		max_heapify(A, largest)
	}
}

// Return the maximum element from the heap (root)
func max_heap_maximum(A heap) int {
	if A.heap_size < 1 {
		panic("heap underflow")
	}
	return A.heap[0] // Max element is the root
}

// Max-Heap Extract-Max function
func max_heap_extract_max(A *heap) int {
	// Get the maximum element
	max := max_heap_maximum(*A)

	// Move the last element to the root and decrease heap size
	A.heap[0] = A.heap[A.heap_size-1]
	A.heap_size--

	// Restore the heap property by heapifying from the root
	max_heapify(A, 0)

	// Return the extracted max
	return max
}

func max_heap_increase_key(A *heap, index int, key int) error {
	if key < A.heap[index] {
		return errors.New("new key is smaller than current key")
	}
	A.heap[index] = key
	for index > 0 && A.heap[parent(index)] < A.heap[index] {
		A.heap[index], A.heap[parent(index)] = A.heap[parent(index)], A.heap[index]
		index = parent(index)
	}
	return nil
}

// MAX-HEAP-INSERT function
func max_heap_insert(A *heap, x int, maxSize int) error {
	if A.heap_size == maxSize {
		return errors.New("heap overflow")
	}

	// Expand the heap and set initial value to -inf
	A.heap_size++
	A.heap = append(A.heap, math.MinInt32)

	// Increase the key to the desired value
	return max_heap_increase_key(A, A.heap_size-1, x)
}

func main() {
	// Example usage
	A := &heap{[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}, 10}
	fmt.Println("Heap before extraction:", A.heap[:A.heap_size])

	max := max_heap_extract_max(A)
	fmt.Println("Extracted max:", max)
	fmt.Println("Heap after extraction:", A.heap[:A.heap_size])

	// Create a new heap and define max size
	A = &heap{heap: []int{}, heap_size: 0}
	maxSize := 10

	// Insert elements
	for _, val := range []int{16, 14, 10, 8, 7, 9} {
		err := max_heap_insert(A, val, maxSize)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	// Display the max heap after insertion
	fmt.Println("Max Heap after insertions:", A.heap[:A.heap_size])

	// Increase the key of the element at index 3
	err := max_heap_increase_key(A, 3, 15)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Max Heap after key increase:", A.heap[:A.heap_size])
	}	
}
