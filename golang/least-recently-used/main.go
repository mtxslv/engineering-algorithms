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
	N := 10 // or 100
	K := 6 // or 3
	utils.SimulationLRUOPT(6,10,N,K)
} 
