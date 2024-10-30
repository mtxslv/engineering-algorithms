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
	TEMPOS_RANDOM_SELECT []int64 `json:"tempos_random_select"` // Capitalized field name for export
	TEMPOS_SELECT []int64 `json:"tempos_select"` // Capitalized field name for export
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write_array_json(arr_n []int32, arr_random_select, arr_select []int64) {
	tempos := Dados{arr_n, arr_random_select, arr_select}
	b, err := json.Marshal(tempos)
	check(err)
	err = os.WriteFile("./comparison-select-shuffled.json", b, 0644)
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
	return arr_copy
}


func partition(A []int, p int, r int) int {
	x := A[r] // the pivot
	i := p - 1 // highest index into the low side
	for j := p; j < r; j++ {
		if A[j] <= x {
			i = i + 1
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

func randomizedPartition(A []int, p int, r int, randomizer *rand.Rand) int {
	i := p + randomizer.Intn(r-p+1)
	A[i], A[r] = A[r], A[i]
	return partition(A, p, r)
}

func randomizedSelect(A []int, p, r, i int, randomizer *rand.Rand) int {
	// 1 <= i <= r-p+1 when p==r means that i == 1
	if p == r {
		return A[p]
	}
	q := randomizedPartition(A, p, r, randomizer)
	k := q - p + 1
	if i == k {
		return A[q] // the pivot value is the answer
	} else if i < k {
		return randomizedSelect(A, p, q-1, i, randomizer)
	} else {
		return randomizedSelect(A, q+1, r, i-k, randomizer)
	}
}


func randomizedQuicksort(A []int, p int, r int, randomizer *rand.Rand) {
	if p < r {
		q := randomizedPartition(A, p, r, randomizer)
		randomizedQuicksort(A, p, q-1, randomizer)
		randomizedQuicksort(A, q+1, r, randomizer)
	}
}

func selectHelper(A []int, p, r, i int, randomizer *rand.Rand) int {
	for (r - p + 1) % 5 != 0 {
		for j := p + 1; j <= r; j++ {
			// put the minimum into A[p]
			if A[p] > A[j] {
				A[p], A[j] = A[j], A[p]
			}
		}
		// If we want the minimum of A[p:r], we're done.
		if i == 1 {
			return A[p]
		}
		p++
		i--
	}

	g := (r - p + 1) / 5 // number of 5-element groups

	// Sort each group of five in place
	// Compare the following loop with
	// for j = p to p + g - 1
	for j := 0; j < g; j++ { 
		group := []int{A[p+j], A[p+j+g], A[p+j+2*g], A[p+j+3*g], A[p+j+4*g]}
		randomizedQuicksort(group,0,len(group)-1,randomizer)
		A[p+j+2*g] = group[2] // Median of group placed in A[p+j+2*g]
	}

	// Recursively find the pivot x as the median of the group medians
	x := selectHelper(A, p+2*g, p+3*g-1, int(math.Ceil(float64(g)/2.0)), randomizer)
	q := partitionAroundPivot(A, p, r, x)

	k := q - p + 1
	if i == k {
		return A[q]
	} else if i < k {
		return selectHelper(A, p, q-1, i, randomizer)
	} else {
		return selectHelper(A, q+1, r, i-k, randomizer)
	}
}

func partitionAroundPivot(A []int, p, r, pivot int) int {
	for i := p; i <= r; i++ {
		if A[i] == pivot {
			A[i], A[r] = A[r], A[i]
			break
		}
	}
	return partition(A, p, r)
}

func main() {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	A := []int{6, 19, 4, 12, 14, 9, 15, 7, 8, 11, 3, 13, 2, 5, 10}
	p := 0              // first element
	r := len(A) - 1     // last index of A
	i := 5              // we want the 5th smallest element
	selection := randomizedSelect(A, p, r, i, randomizer)
	result := selectHelper(A, 0, r, i, randomizer)
	fmt.Printf("The value is: %d \n(Must be 6)\n", selection)
	fmt.Printf("The value is: %d \n(Must be 6)\n", result)
}
