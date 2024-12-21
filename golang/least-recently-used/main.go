// package main

// import "utils/utils"

// func main() {
// 	utils.SimulationWithRandomness()
// }

package main

import (
	"utils/utils"
)

func main(){
	universeValues := 5000
	codeSize := 10
	N := 500
	for N < 1000 {
		K := 3 
		for K < 15 {
			utils.SimulationLRUOPT(codeSize,universeValues,N,K)
			K++
		}
		N++
	}
} 
