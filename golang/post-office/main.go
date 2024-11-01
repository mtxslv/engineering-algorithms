package main

import (
	"fmt"
	"math/rand/v2"
)

type weightedPoint struct {
	point float64
	weight float64
}

// QUICKSORT

func partition(A []weightedPoint, p int, r int) int {
	x := A[r].point // the pivot
	i := p - 1 // highest index into the low side
	for j := p ; j < r ; j++ {
		if A[j].point <= x {
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

func randomized_partition(A []weightedPoint, p int, r int, randomizer *rand.Rand) int {
	i := int32(p) + randomizer.Int32N(int32(r-p))
	aux := A[i]
	A[i] = A[r]
	A[r] = aux
	return partition(A,p,r)
}

func randomizedQuicksort(A []weightedPoint, p int, r int, randomizer *rand.Rand){
	if p < r {
		// Partition the subarray around the pivot, which ends up in A[q].
		q := randomized_partition(A,p,r, randomizer)
		randomizedQuicksort(A, p, q-1, randomizer) // recursively sort the low side
		randomizedQuicksort(A, q+1, r, randomizer) // recursively sort the high side
	}
}

// MAIN

func weightedMedian(A []weightedPoint) weightedPoint {
	randomizer := rand.New(rand.NewPCG(255, 16515616))
	// Step 1: order the array
	randomizedQuicksort(A, 0, len(A)-1, randomizer)
	
	// Step 2: Calculate the total weight
	totalWeight := 0.0
	for _, point := range A {
		totalWeight += point.weight
	}

	ans := weightedPoint{0.0,0.0} // placeholder
	// Step 3: Find the weighted median
	cumulativeWeight := 0.0
	for _, point := range A {
		cumulativeWeight += point.weight
		if cumulativeWeight >= totalWeight/2 {
			ans = point
			break
		}
	}
	return ans
}


func main(){
	A_x := []weightedPoint{
		{ 1.0, 0.075,},
		{ 8.0, 0.35,},
		{ 6.0, 0.2,},
		{ 5.0, 0.08,},
		{ 4.0, 0.15,},
		{ 3.0, 0.12,},
		{ 2.0, 0.025,},
	}
	A_y := []weightedPoint{
		{ -3.0, 0.075,},
		{ 4.0, 0.15,},
		{  0.0, 0.12,},
		{ 18.0, 0.35,},
		{ 7.0, 0.2,},
		{ -2.0, 0.025,},
		{ 5.0, 0.08,},
	}	
	x_ans := weightedMedian(A_x)
	y_ans := weightedMedian(A_y)
	fmt.Printf("Point where minimum occurs: (%.1f, %.1f)",x_ans.point, y_ans.point)
}