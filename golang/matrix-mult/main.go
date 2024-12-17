package main

import (
	"fmt"
	"utils/utils"
)

func main(){
	A := [][]float32{
		{1,2,3},
		{4,5,6},
	}

	B := [][]float32{
		{7,8},
		{9,10},
		{11,12},
	}

	var C [][]float32

	C, costs := utils.MatMultWithCosts(&A,&B)

	for _, cost := range costs {
		fmt.Printf("%s\n", cost)
	}
	
	fmt.Printf("\n")
	for _, row := range C {
		fmt.Printf("%+v\n", row)
	}
}