package main

import "utils/utils"

func main() {
	requests := []int{1, 2, 3, 4, 1, 2, 5, 1, 2, 3, 4, 5}
	for cacheSize := 2; cacheSize <= 5; cacheSize++ {
		utils.SimulateRequests(cacheSize, requests)
	}
}


// package main

// import "fmt"

// func main(){
// 	var cacheHashTable = map[int]string{}
// 	cacheHashTable[1] = "baguga"
// 	cacheHashTable[2] = "glub glub"
// 	cacheHashTable[3] = "mamada"
// 	cacheHashTable[4] = ""
	
// 	for k := range cacheHashTable{
// 		fmt.Printf("%s\n",cacheHashTable[k])
// 	}

// 	delete(cacheHashTable,2)

// 	fmt.Printf("\n\nDELETED MIDDLE ITEM\n\n")

// 	for k := range cacheHashTable{
// 		fmt.Printf("%s\n",cacheHashTable[k])
// 	}
// }