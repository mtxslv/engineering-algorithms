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

func copyArray(arr []int) []int { 
	arr_copy := make([]int, len(arr))
	for i := 0; i < len(arr); i++{
		arr_copy[i] = arr[i]
	}	
	are_equal := compare_array(arr,arr_copy)
	if !are_equal{
		fmt.Println("Sorted arrays differ. Stopping...")
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

func selectSmartPivot(A []int, p, r, i int, randomizer *rand.Rand) int {
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
	x := selectSmartPivot(A, p+2*g, p+3*g-1, int(math.Ceil(float64(g)/2.0)), randomizer)
	q := partitionAroundPivot(A, p, r, x)

	k := q - p + 1
	if i == k {
		return A[q]
	} else if i < k {
		return selectSmartPivot(A, p, q-1, i, randomizer)
	} else {
		return selectSmartPivot(A, q+1, r, i-k, randomizer)
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
	MIN := 10
	MAX :=  60000
	
	i := 0
	
	times_select := make([]int64, MAX - MIN + 1)
	times_random_select := make([]int64, MAX - MIN + 1)
	arr_n := make([]int32, MAX - MIN + 1)

	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	for n := MIN ; n <= MAX; n++ {
		// Generate shuffled array
		A_select := shuffledArray(n, randomizer)
		A_random_select := copyArray(A_select)

		fmt.Println("ARRAY: ", A_random_select)
		fmt.Println("ARRAY: ", A_select)
		
		order_stats := 5              // we want the 5th smallest element

		// Search for values
		start_random_select := time.Now()
		randomized_ans := randomizedSelect(A_random_select, 0, len(A_random_select) - 1, order_stats, randomizer)
		random_select_time := time.Since(start_random_select)

		start_select := time.Now()
		select_ans := selectSmartPivot(A_select, 0, len(A_select) - 1, order_stats, randomizer)
		select_time := time.Since(start_select)

		fmt.Printf("Order Stats: %d (Random) | %d (Smart Pivot)\n", randomized_ans, select_ans)

		if select_ans != randomized_ans {
			fmt.Println("Order Statistics found differ. Stopping...")
			break
		} 
			
		times_random_select[i] =  random_select_time.Nanoseconds()
		times_select[i] = select_time.Nanoseconds()
		arr_n[i] = int32(n)
		i = i+1		
		
		fmt.Printf("Run %d finished\n", n)
		
	}
	fmt.Print("Salvando Arquivo...")
	write_array_json(arr_n, times_random_select, times_select)
}	