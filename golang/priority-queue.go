package main

import "fmt"

type heap struct {
	heap []int
	heap_size int
}

func parent(i int) int  { return (i - 1) / 2 }
func left(i int) int    { return 2*i + 1 }
func right(i int) int   { return 2*i + 2 }

func max_heapify(A heap, i int) {
    var l = left(i)
    var r = right(i)
    var largest = i // Initially, assume 'i' is the largest

    // Ensure 'l' is within bounds and check if left child is larger
    if l < A.heap_size && A.heap[l] > A.heap[i] {
        largest = l
    }
    // Ensure 'r' is within bounds and check if right child is larger
    if r < A.heap_size && A.heap[r] > A.heap[largest] {
        largest = r
    }

    // If 'largest' is not 'i', swap and continue heapifying
    if largest != i {
		aux := A.heap[i]
		A.heap[i] = A.heap[largest]
		A.heap[largest] = aux
        max_heapify(A, largest)
    }
}

func build_max_heap(A *heap, n int){
	A.heap_size = n
	for i := n/2; i >= 0; i-- {
		max_heapify(*A,i)
	}
}

func max_heap_maximum(A heap) int {
	if A.heap_size < 1{
		return -1
	}else{
		return A.heap[0]
	}
}

func max_heap_extract_max(A heap) int {
	max := max_heap_maximum(A)
	A.heap[0] = A.heap[A.heap_size]
	A.heap_size = A.heap_size - 1
	max_heapify(A,0)
	return max
}

func main(){
	var expected_heap = heap{[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1,},10} 
	a := max_heap_maximum(expected_heap)
	fmt.Println(a)
}