package main

import (
	"fmt"
)

// A utility function to get the maximum value in an array
func getMax(A []int) int {
	max := A[0]
	for _, num := range A {
		if num > max {
			max = num
		}
	}
	return max
}

// Counting Sort that sorts A based on the digit represented by exp
func countingSortByDigit(A []int, n, exp int) []int {
	B := make([]int, n)    // Output array
	C := make([]int, 10)   // There are 10 possible digits (0-9)

	// Store count of occurrences in C
	for i := 0; i < n; i++ {
		index := (A[i] / exp) % 10
		C[index]++
	}

	// Change C[i] so that C[i] contains the actual position of the digit in B
	for i := 1; i < 10; i++ {
		C[i] += C[i-1]
	}

	// Build the output array B
	for i := n - 1; i >= 0; i-- {
		index := (A[i] / exp) % 10
		B[C[index]-1] = A[i]
		C[index]--
	}

	// Copy the output array B back to A
	for i := 0; i < n; i++ {
		A[i] = B[i]
	}

	return A
}

// Radix Sort function
func radixSort(A []int, n int) []int {
	// Find the maximum number to determine the number of digits
	max := getMax(A)

	// Apply counting sort to each digit
	// exp is 10^i where i is the current digit place (1s, 10s, 100s, etc.)
	for exp := 1; max/exp > 0; exp *= 10 {
		A = countingSortByDigit(A, n, exp)
	}

	return A
}

func main() {
	// Example usage
	A := []int{170, 45, 75, 90, 802, 24, 2, 66}
	n := len(A)

	fmt.Println("Original array:", A)
	sortedArray := radixSort(A, n)
	fmt.Println("Sorted array:", sortedArray)
}
