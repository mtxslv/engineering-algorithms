package utils

import (
	// "errors"
	"fmt"
)

func getMatrixDimension(pM *[][]float32) (int,int,error){
	M := *pM
	rows := len(M)
	cols := len(M[0])
	for _, row := range M {
		if len(row) != cols {
			return -1,-1,ErrMalformedMatrix
		}
	}
	return rows, cols, nil
}

func matMult(pointerA, pointerB *[][]float32) [][]float32 {
	// check dimensions of A and B
	rowsA, colsA, errA := getMatrixDimension(pointerA)
	rowsB, colsB, errB := getMatrixDimension(pointerB)

	if errA!=nil{
		panic(errA)
	}
	if errB != nil {
		panic(errB)
	}
	
	// check if A and B can be multiplied
	if colsA != rowsB {
		panic(ErrBadDimensions)
	}

	// Dereference
	A := *pointerA
	B := *pointerB

	var C [][]float32

	// Create C
	i := 0
	for i < rowsA {
		C = append(C, make([]float32,colsB))
		i++
	}

	var itRow, itCol, itK int

	for itRow < rowsA{
		itCol = 0
		for itCol < colsB {
			itK = 0
			for itK < rowsB {
				C[itRow][itCol] += A[itRow][itK]*B[itK][itCol] 
				itK++
			}
			itCol++
		}
		itRow++
	}

	return C
}


func MatMultWithCosts(pointerA, pointerB *[][]float32) ([][]float32,[]string) {
	
	// Cost variable
	var costs []string
	
	// check dimensions of A and B
	rowsA, colsA, errA := getMatrixDimension(pointerA)
	rowsB, colsB, errB := getMatrixDimension(pointerB)

	if errA!=nil{
		panic(errA)
	}
	if errB != nil {
		panic(errB)
	}
	
	// check if A and B can be multiplied
	if colsA != rowsB {
		panic(ErrBadDimensions)
	}

	// Dereference
	A := *pointerA
	B := *pointerB

	var C [][]float32
	
	// Create C
	// aGivenRow := make([]float32, colsB)
	i := 0
	for i < rowsA {
		C = append(C, make([]float32,colsB))
		i++
	}

	var itRow, itCol, itK int

	for itRow < rowsA{
		itCol = 0
		costs = append(costs, "itRow")
		costs = append(costs, "rowsA")
		costs = append(costs, "itCol")
		for itCol < colsB {
			itK = 0
			costs = append(costs, "itCol")
			costs = append(costs, "colsB")
			costs = append(costs, "itK")
			for itK < rowsB {
				costs = append(costs,"rowsB")
				costs = append(costs, "itK")
				costs = append(costs, "itCol")
				costs = append(costs, "itRow")
				costs = append(costs,fmt.Sprintf("A_%d_%d",itRow,itK))
				costs = append(costs,fmt.Sprintf("B_%d_%d",itK,itCol))
				costs = append(costs,fmt.Sprintf("C_%d_%d",itRow,itCol))
				costs = append(costs, "itK")
				C[itRow][itCol] += A[itRow][itK]*B[itK][itCol] 
				itK++
			}
			costs = append(costs, "itCol")
			itCol++
		}
		costs = append(costs, "itRow")
		itRow++
	}

	return C, costs
}