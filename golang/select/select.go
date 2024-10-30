package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func partition(A []int, p int, r int) int {
	x := A[r] // the pivot
	i := p - 1 // highest index into the low side
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

func randomized_partition(A []int, p int, r int, randomizer *rand.Rand) int {
	i := p + int(randomizer.Int31n(int32(r-p+1)))
	A[i], A[r] = A[r], A[i]
	return partition(A, p, r)
}

func randomizedQuicksort(A []int, p int, r int, randomizer *rand.Rand) {
	if p < r {
		q := randomized_partition(A, p, r, randomizer)
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

func SELECT(A []int, i int) int {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	return selectHelper(A, 0, len(A)-1, i, randomizer)
}

func main() {
	A := []int{12, 3, 5, 7, 4, 19, 26}
	i := 3 // Find the 3rd smallest element
	result := SELECT(A, i)
	fmt.Printf("The %d-th smallest element is %d\n", i, result)
}
