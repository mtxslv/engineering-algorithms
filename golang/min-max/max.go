package main

import (
	"fmt"
)

func findMax(arr []int, n int) int {

	max := arr[1]
	for i := 1; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	} 

	return max
}


func main() {
	arr := []int{3, 5, 1, 2, 4, 8, 6}
	n := len(arr)
	max := findMax(arr, n)
	fmt.Printf("Maximum: %d\n", max)
}
