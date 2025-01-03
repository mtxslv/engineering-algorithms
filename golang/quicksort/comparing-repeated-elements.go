package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"	
    "time"
)

// JSON STUFF

type Dados struct {
	PERC []float32 `json:"n"` // Capitalized field name for export
	TEMPOS_RANDOMIZED []int64 `json:"tempos_randomized"` // Capitalized field name for export
	TEMPOS_REGULAR []int64 `json:"tempos_regular"` // Capitalized field name for export
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write_array_json(arr_perc []float32, arr_randomized, arr_regular []int64) {
	tempos := Dados{arr_perc, arr_randomized, arr_regular}
	b, err := json.Marshal(tempos)
	check(err)
	err = os.WriteFile("./comparison-repeated-quicksort-and-randomized.json", b, 0644)
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

func repeat_elements(p float32, n int, arr []int) []int {
	how_many := int(p*float32(n))
	fmt.Print("REPEAT ", how_many)
	for i := 0; i < how_many; i++{
		arr[n-i-1] = arr[0]
	}
	return arr
}

func display_array(n int, arr []int) {
	msg := fmt.Sprintf("GENERATED ARRAY (%d): \n %d", n, arr)
	fmt.Println(msg)
}

func shuffled_array(n int, r *rand.Rand) []int {
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

// QUICKSORT

func partition(A []int, p int, r int) int {
	x := A[r] // the pivot
	i := p - 1 // highest index into the low side
	for j := p ; j < r ; j++ {
		if A[j] <= x {
			i = i + 1
			aux := A[i]
			A[i] = A[j]
			A[j] = aux
		}
	} 
	aux := A[r]
	A[r] = A[i+1]
	A[i+1] = aux
	return i + 1
}

func quicksort(A []int, p int, r int){
	if p < r {
		// Partition the subarray around the pivot, which ends up in A[q].
		q := partition(A,p,r)
		quicksort(A, p, q-1) // recursively sort the low side
		quicksort(A, q+1, r) // recursively sort the high side
	}
}

// RANDOMIZED QUICKSORT

func randomized_partition(A []int, p int, r int, randomizer *rand.Rand) int {
	i := int32(p) + randomizer.Int32N(int32(r-p))
	aux := A[i]
	A[i] = A[r]
	A[r] = aux
	return partition(A,p,r)
}

func randomized_quicksort(A []int, p int, r int, randomizer *rand.Rand){
	if p < r {
		// Partition the subarray around the pivot, which ends up in A[q].
		q := randomized_partition(A,p,r, randomizer)
		randomized_quicksort(A, p, q-1, randomizer) // recursively sort the low side
		randomized_quicksort(A, q+1, r, randomizer) // recursively sort the high side
	}
}

// MAIN

func main(){
	N := 40000 // 3000
	
	arr_perc := []float32{.1, .15, .2, .25, .3, .35, .4, .45, .5, .55, .6, .65, .7, .75, .8, .85, .9}
	times_quicksort := make([]int64, len(arr_perc))
	times_randomized := make([]int64, len(arr_perc))

	r := rand.New(rand.NewPCG(255, 16515616))
	
	for i := 0 ; i < len(arr_perc); i++ {
		p := arr_perc[i]
		fmt.Println("Running for p = ", p)
		arr_regular := shuffled_array(N, r)
		arr_randomized := copyArray(arr_regular)

		start_quicksort := time.Now()
		quicksort(arr_regular,0,N-1)
		regular_quicksort_time := time.Since(start_quicksort)

		start_randomized := time.Now()
		randomized_quicksort(arr_randomized,0,N-1,r)
		randomized_quicksort_time := time.Since(start_randomized)

		are_equal := compare_array(arr_regular,arr_randomized)
		if !are_equal{
			fmt.Println("Sorted arrays differ. Stopping...")
			break
		}

		times_quicksort[i] =  regular_quicksort_time.Nanoseconds()
		times_randomized[i] = randomized_quicksort_time.Nanoseconds()
	}
	fmt.Print("Salvando Arquivo...")
	write_array_json(arr_perc, times_randomized, times_quicksort)
}