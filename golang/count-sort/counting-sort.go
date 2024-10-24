package main

import (
	"fmt"
)

// ARRAY GENERATION / DISPLAY

func generate_array(n int, reversed bool) []int {
	arr := make([]int, n)
	for i := 0 ; i < n; i++ {
		if reversed {
			arr[i] = n - i - 1
		} else {
			arr[i] = i
		}
	} 
	return arr
}

func countingsort(A []int, n, k int) []int {
	B := make([]int, n)   // B array with size n
	C := make([]int, k+1) // C array with size k+1

	// Count occurrences of each element in A
	for j := 0; j < n; j++ {
		C[A[j]] = C[A[j]] + 1
	}

	// C[i] now contains the number of elements less than or equal to i
	for i := 1; i <= k; i++ {
		C[i] = C[i] + C[i-1]
	}

	// Build the sorted array B
	for j := n - 1; j >= 0; j-- {
		B[C[A[j]]-1] = A[j]
		C[A[j]] = C[A[j]] - 1
	}

	return B
}

func main() {
	A := []int{4, 2, 2, 8, 3, 3, 1, 7}
	k := 8 // Maximum value in A
	n := len(A)

	fmt.Println("Original array:", A)
	sortedArray := countingsort(A, n, k)
	fmt.Println("Sorted array:", sortedArray)
}