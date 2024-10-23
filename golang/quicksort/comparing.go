package main

import (
	"fmt"
	"math/rand/v2"
    "time"	
)

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

func display_array(n int, arr []int) {
	msg := fmt.Sprintf("GENERATED ARRAY (%d): \n %d", n, arr)
	fmt.Println(msg)
}

func shuffled_array(n int, r *rand.Rand) []int {
	return r.Perm(n)
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
	MIN := 10
	MAX := 1750
	
	i := 0
	
	times_quicksort := make([]int64, MAX - MIN + 1)
	times_randomized := make([]int64, MAX - MIN + 1)
	arr_n := make([]int32, MAX - MIN + 1)

	r := rand.New(rand.NewPCG(255, 16515616))
	for n := MIN ; n <= MAX; n++ {
		fmt.Println("Running for n = ", n)
		arr_regular := generate_array(n,true)
		arr_randomized := generate_array(n,true)

		start_quicksort := time.Now()
		quicksort(arr_regular,0,n-1)
		regular_quicksort_time := time.Since(start_quicksort)

		start_randomized := time.Now()
		randomized_quicksort(arr_randomized,0,n-1,r)
		randomized_quicksort_time := time.Since(start_randomized)

		are_equal := compare_array(arr_regular,arr_randomized)
		if !are_equal{
			fmt.Println("Sorted arrays differ. Stopping...")
			break
		}

		times_quicksort[i] =  regular_quicksort_time.Nanoseconds()
		times_randomized[i] = randomized_quicksort_time.Nanoseconds()
		arr_n[i] = int32(n)
		i = i+1
	}
	fmt.Print(times_quicksort)
}