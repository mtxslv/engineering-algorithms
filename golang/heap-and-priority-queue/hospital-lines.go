package main

import (
	"fmt"
	"math"
	"errors"
)

type Ticket struct {
	number int
	risk   int
}

type heap struct {
	heap      []Ticket
	heap_size int
}

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return 2*i + 1 }
func right(i int) int  { return 2*i + 2 }

// Max heapify function
func max_heapify(A *heap, i int) {
	l := left(i)
	r := right(i)
	largest := i // Initially, assume 'i' is the largest

	// Ensure 'l' is within bounds and check if left child has a higher risk
	if l < A.heap_size && A.heap[l].risk > A.heap[i].risk {
		largest = l
	}
	// Ensure 'r' is within bounds and check if right child has a higher risk
	if r < A.heap_size && A.heap[r].risk > A.heap[largest].risk {
		largest = r
	}

	// If 'largest' is not 'i', swap and continue heapifying
	if largest != i {
		A.heap[i], A.heap[largest] = A.heap[largest], A.heap[i]
		max_heapify(A, largest)
	}
}

// Return the maximum element from the heap (root)
func max_heap_maximum(A heap) Ticket {
	if A.heap_size < 1 {
		panic("heap underflow")
	}
	return A.heap[0] // Max element is the root
}

// Max-Heap Extract-Max function
func max_heap_extract_max(A *heap) Ticket {
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

func max_heap_increase_key(A *heap, index int, risk int) error {
	if risk < A.heap[index].risk {
		return errors.New("new risk is smaller than current risk")
	}
	A.heap[index].risk = risk
	for index > 0 && A.heap[parent(index)].risk < A.heap[index].risk {
		A.heap[index], A.heap[parent(index)] = A.heap[parent(index)], A.heap[index]
		index = parent(index)
	}
	return nil
}

// MAX-HEAP-INSERT function
func max_heap_insert(A *heap, x Ticket, maxSize int) error {
	if A.heap_size == maxSize {
		return errors.New("heap overflow")
	}

	// Expand the heap and set initial value to the lowest possible risk
	A.heap_size++
	A.heap = append(A.heap, Ticket{number: x.number, risk: math.MinInt32})

	// Increase the key to the desired risk
	return max_heap_increase_key(A, A.heap_size-1, x.risk)
}

func examinePatient(A *heap){
	patientNumber := len(A.heap)
	for i := 0; i < patientNumber; i++ {
		ticket := max_heap_extract_max(A)
		fmt.Printf("Ticket number: %d, risk: %d\n", ticket.number, ticket.risk)
	}
}

func main() {
	// Create a new heap and define max size
	A := &heap{heap: []Ticket{}, heap_size: 0}
	maxSize := 10

	// Insert elements
	tickets := []Ticket{
		{number: 1,  risk: 16},
		{number: 2,  risk: 4},
		{number: 3,  risk: 8},
		{number: 4,  risk: 14},
		{number: 5,  risk: 1},
		{number: 6,  risk: 9},
		{number: 7,  risk: 3},
		{number: 8,  risk: 2},
		{number: 9,  risk: 10},
		{number: 10, risk: 7},		
	}

	for _, ticket := range tickets {
		err := max_heap_insert(A, ticket, maxSize)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	examinePatient(A)

	// // Display the max heap after insertion
	// fmt.Println("Max Heap after insertions:")
	// for _, ticket := range A.heap[:A.heap_size] {
	// 	fmt.Printf("Ticket number: %d, risk: %d\n", ticket.number, ticket.risk)
	// }

	// // Increase the risk of the ticket at index 3
	// err := max_heap_increase_key(A, 3, 15)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("\nMax Heap after risk increase:")
	// 	for _, ticket := range A.heap[:A.heap_size] {
	// 		fmt.Printf("Ticket number: %d, risk: %d\n", ticket.number, ticket.risk)
	// 	}
	// }
}
