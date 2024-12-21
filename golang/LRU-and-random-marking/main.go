package main

import (
	"fmt"
	"os"
	"utils/utils"
)

func getData() map[string]float32 {
	data := make(map[string]float32)

	data["A_0_0"] = 0 
	data["A_0_1"] = 0
	data["A_1_0"] = 0
	data["A_1_1"] = 0
	data["A_2_0"] = 0
	data["A_2_1"] = 0
	data["A_3_0"] = 0
	data["A_3_1"] = 0

	data["B_0_0"] = 0
	data["B_0_1"] = 0
	data["B_0_2"] = 0
	data["B_0_3"] = 0
	data["B_1_0"] = 0
	data["B_1_1"] = 0
	data["B_1_2"] = 0
	data["B_1_3"] = 0

	data["C_0_0"] = 0
	data["C_0_1"] = 0
	data["C_0_2"] = 0
	data["C_0_3"] = 0
	data["C_1_0"] = 0
	data["C_1_1"] = 0
	data["C_1_2"] = 0
	data["C_1_3"] = 0
	data["C_2_0"] = 0
	data["C_2_1"] = 0
	data["C_2_2"] = 0
	data["C_2_3"] = 0
	data["C_3_0"] = 0
	data["C_3_1"] = 0
	data["C_3_2"] = 0
	data["C_3_3"] = 0

	data["colsB"] = 0
	data["itCol"] = 0
	data["itK"] = 0
	data["itRow"] = 0
	data["rowsA"] = 0
	data["rowsB"] = 0
	return data
}

func main(){
	// capacity := 10
	// lru := utils.NewLRUCacheV1(capacity)
	// rmc, err := utils.NewRandomMarkingCache(capacity)
	// if err != nil {
	// 	panic(err)
	// }

	// Check if parameters were provided
	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		fmt.Printf("Usage: \n./main <path-to-txt-file>\n")
		return
	} 

	filePath := argsWithProg[1]

	txt := utils.LoadLines(filePath)

	
	// fmt.Printf("%d\n",len(txt))
	for it, line := range txt {
		// fmt.Printf("WHAT SO\n")
		fmt.Printf("LINE %d : %s\n", it, line)
	}

}