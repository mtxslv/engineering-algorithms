package main

import (
	"fmt"
	"math"
)

// findMinMax function finds the minimum and maximum values in an array
func findMinMax(arr []int) (int, int) {
	n := len(arr)

	if n == 0 {
		return math.MaxInt, math.MinInt // Return Max and Min values if array is empty
	}

	var min, max int
	if n%2 == 0 { // Even number of elements
		if arr[0] < arr[1] {
			min, max = arr[0], arr[1]
		} else {
			min, max = arr[1], arr[0]
		}
		startIndex := 2

		for i := startIndex; i < n; i += 2 {
			if arr[i] < arr[i+1] {
				min = minInt(min, arr[i])
				max = maxInt(max, arr[i+1])
			} else {
				min = minInt(min, arr[i+1])
				max = maxInt(max, arr[i])
			}
		}
	} else { // Odd number of elements
		min, max = arr[0], arr[0]
		startIndex := 1

		for i := startIndex; i < n; i += 2 {
			if i+1 < n { // If there's a pair
				if arr[i] < arr[i+1] {
					min = minInt(min, arr[i])
					max = maxInt(max, arr[i+1])
				} else {
					min = minInt(min, arr[i+1])
					max = maxInt(max, arr[i])
				}
			} else { // Handle the case for an odd element
				min = minInt(min, arr[i])
			}
		}
	}

	return min, max
}

// minInt returns the minimum of two integers
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// maxInt returns the maximum of two integers
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMax(arr []int, n int) int {

	max := arr[1]
	for i := 1; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	} 

	return max
}

func findMin(arr []int, n int) int {

	min := arr[1]
	for i := 1; i < n; i++ {
		if min > arr[i] {
			min = arr[i]
		}
	} 

	return min
}

func main() {
	arr := []int{3, 5, 1, 2, 4, 8, 6}
	min, max := findMinMax(arr)
	fmt.Printf("Minimum: %d, Maximum: %d\n", min, max)
}
