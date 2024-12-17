package main

import (
	"fmt"
	"utils/utils"
)

func main(){
	A := [][]float32{
		{1,2},
		{2,4},
		{5,6},
		{7,8},
	}

	B := [][]float32{
		{5,1,8,1},
		{9,2,5,9},
	}

	_, costs := utils.MatMultWithCosts(&A,&B)

	for _, cost := range costs {
		fmt.Printf("%s\n", cost)
	}
	
	// fmt.Printf("\n")
	// for _, row := range C {
	// 	fmt.Printf("%+v\n", row)
	// }
}