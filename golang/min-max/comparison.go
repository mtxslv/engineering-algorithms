package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"	
    "time"
)

// JSON STUFF

type Dados struct {
	N []int32 `json:"n"` // Capitalized field name for export
	TEMPOS_SIMULTANEOUS []int64 `json:"tempos_simultaneous"` // Capitalized field name for export
	TEMPOS_STANDALONE []int64 `json:"tempos_standalone"` // Capitalized field name for export
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write_array_json(arr_n []int32, arr_simultaneous, arr_standalone []int64) {
	tempos := Dados{arr_n, arr_simultaneous, arr_standalone}
	b, err := json.Marshal(tempos)
	check(err)
	err = os.WriteFile("./comparison-minmax.json", b, 0644)
	check(err)
}

func shuffledArray(n int, r *rand.Rand) []int {
	return r.Perm(n)
}

func copyArray(arr []int) []int { 
	arr_copy := make([]int, len(arr))
	for i := 0; i < len(arr); i++{
		arr_copy[i] = arr[i]
	}
	// are_equal := compare_array(arr,arr_copy)
	// if !are_equal{
	// 	fmt.Println("Sorted arrays differ. Stopping...")
	// }	
	return arr_copy
}


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

	max := arr[0]
	for i := 1; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	} 

	return max
}

func findMin(arr []int, n int) int {

	min := arr[0]
	for i := 1; i < n; i++ {
		if min > arr[i] {
			min = arr[i]
		}
	} 

	return min
}

func main() {
	MIN := 10
	MAX :=  60000
	
	i := 0
	
	times_standalone := make([]int64, MAX - MIN + 1)
	times_simultaneous := make([]int64, MAX - MIN + 1)
	arr_n := make([]int32, MAX - MIN + 1)

	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	for n := MIN ; n <= MAX; n++ {
		// Generate shuffled array
		arr := shuffledArray(n, randomizer)

		// Search for values
		start_simultaneous := time.Now()
		min_simultaneous, max_simultaneous := findMinMax(arr)
		simultaneous_time := time.Since(start_simultaneous)

		start_standalone := time.Now()
		min_standalone := findMin(arr, len(arr))
		max_standalone := findMax(arr, len(arr))
		standalone_time := time.Since(start_standalone)
		
		if min_simultaneous != min_standalone || max_simultaneous != max_standalone {
			fmt.Println("Min/Max found differ. Stopping...")
			break
		} 
			
		times_simultaneous[i] =  simultaneous_time.Nanoseconds()
		times_standalone[i] = standalone_time.Nanoseconds()
		arr_n[i] = int32(n)
		i = i+1		
		
		fmt.Printf("Run %d finished\n", n)
		
	}
	fmt.Print("Salvando Arquivo...")
	write_array_json(arr_n, times_simultaneous, times_standalone)
}
