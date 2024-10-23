package main

import (
	"fmt"
	"math/rand/v2"
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

// MAIN

func main(){
	n := 15
	arr := generate_array(n,true)
	fmt.Println("ORIGINAL ARRAY:", arr)
	quicksort(arr,0,n-1)
	fmt.Println("SORTED ARRAY:", arr)
	// r := rand.New(rand.NewPCG(255, 16515616))
	// for n := 1; n <=7 ; n ++ {
	// 	arr := generate_array(n,true)
	// 	// arr := shuffled_array(n,r)
	// 	display_array(n,arr)
	// }
}