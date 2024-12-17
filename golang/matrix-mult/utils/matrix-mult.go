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

func matMult(pointerA, pointerB, pointerC *[][]float32) {
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
	C := *pointerC

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

	fmt.Printf("C = \n")
	for _,row := range C {
		fmt.Printf("%+v\n",row)
	}

}