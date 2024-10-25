package main

import "fmt"

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

func main() {
	// Example usage
	A := &heap{[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}, 10}
	fmt.Println("Heap before extraction:", A.heap[:A.heap_size])

	max := max_heap_extract_max(A)
	fmt.Println("Extracted max:", max)
	fmt.Println("Heap after extraction:", A.heap[:A.heap_size])
}
