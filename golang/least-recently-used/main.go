package main

import "utils/utils"

func main() {
	utils.SimpleSimulation()
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