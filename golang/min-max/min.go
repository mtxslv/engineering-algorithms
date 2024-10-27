package main

import (
	"fmt"
)

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
	n := len(arr)
	min := findMin(arr, n)
	fmt.Printf("Minimum: %d\n", min)
}
