package main

import (
	"fmt"
	"math/rand/v2"
)


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

func randomizedPartition(A []int, p int, r int, randomizer *rand.Rand) int {
	i := int32(p) + randomizer.Int32N(int32(r-p))
	aux := A[i]
	A[i] = A[r]
	A[r] = aux
	return partition(A,p,r)
}

func randomizedSelect(A []int, p, r, i int, randomizer *rand.Rand) int {
	// 1 <= i <= r-p+1 when p==r means that i == 1
	if p == r {
		return A[p] 
	}
	q := randomizedPartition(A,p,r, randomizer)
	k := q - p + 1
	if i == k {
		return A[q] // the pivot value is the answer
	}else{
		if i < k {
			return randomizedSelect(A,p,q-1,i, randomizer)
		} else {
			return randomizedSelect(A,q+1,r,i-k, randomizer)
		}
	}
}

func main() {
	randomizer := rand.New(rand.NewPCG(255, 16515616))
	A := []int{
		6,19,4,12,14,9,15,7,8,11,3,13,2,5,10,
	}
	p := 0 // first element
	r := 14 // last position
	i := 5
	selection := randomizedSelect(A,p,r,i,randomizer)
	fmt.Printf("The value is: %d \n(Must be 6)", selection)
}