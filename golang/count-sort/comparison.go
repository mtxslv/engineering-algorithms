package main

import (
	"encoding/json"
	"fmt"
	"os"	
    "time"
)

// JSON STUFF

type Dados struct {
	N []int32 `json:"n"` // Capitalized field name for export
	TEMPOS_COUNTING []int64 `json:"tempos_counting"` // Capitalized field name for export
	TEMPOS_RADIX []int64 `json:"tempos_radix"` // Capitalized field name for export
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write_array_json(arr_n []int32, arr_counting, arr_radix []int64) {
	tempos := Dados{arr_n, arr_counting, arr_radix}
	b, err := json.Marshal(tempos)
	check(err)
	err = os.WriteFile("./comparison-counting-and-radix.json", b, 0644)
	check(err)
}

// ARRAY GENERATION / DISPLAY / COMPARISON

func compare_array(arr_1 []int, arr_2 []int) bool { // True if arrays equal
	leng_arr_1 := len(arr_1)
	leng_arr_2 := len(arr_2)
	if leng_arr_1 != leng_arr_2{
		return false
	} else {
		for i := 0; i < leng_arr_1; i++ {
			if arr_1[i] != arr_2[i] {
				return false
			}
		}
	}
	return true
}

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

func countingSort(A []int, n, k int) []int {
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
	MIN := 10
	MAX := 70000
	
	i := 0
	
	times_quicksort := make([]int64, MAX - MIN + 1)
	times_randomized := make([]int64, MAX - MIN + 1)
	arr_n := make([]int32, MAX - MIN + 1)

	for n := MIN ; n <= MAX; n++ {
		fmt.Println("Running for n = ", n)
		arr_radix := generate_array(n,true)
		arr_counting := generate_array(n,true)

		start_radix := time.Now()
		ans_radix := radixSort(arr_radix,n)
		radix_time := time.Since(start_radix)

		start_counting := time.Now()
		ans_counting := countingSort(arr_counting, n, n-1)
		counting_time := time.Since(start_counting)

		are_equal := compare_array(ans_radix, ans_counting)
		if !are_equal{
			fmt.Println("Sorted arrays differ. Stopping...")
			break
		}

		times_quicksort[i] =  radix_time.Nanoseconds()
		times_randomized[i] = counting_time.Nanoseconds()
		arr_n[i] = int32(n)
		i = i+1
	}
	fmt.Print("Salvando Arquivo...")
	write_array_json(arr_n, times_randomized, times_quicksort)
}
