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
	N := 500
	for N < 1000 {
		K := 3 
		for K < 15 {
			utils.SimulationLRUOPT(6,10,N,K)
			K++
		}
		N++
	}
} 
