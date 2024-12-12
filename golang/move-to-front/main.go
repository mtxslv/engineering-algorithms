package main

import (
	"fmt"
	"os"
	"strconv"
	"utils/utils"
)

func main() {

	// Check if parameters were provided
	argsWithProg := os.Args
	if len(argsWithProg) != 3 {
		fmt.Printf("Usage: \n./main <list-size> <number-of-requests>\n")
		return
	} 
	
	// If everything is fine, code follows...

	// Define parameters
	listSize, errArg1 := strconv.Atoi(argsWithProg[1])
	requestsNumber, errArg2 := strconv.Atoi(argsWithProg[2])
	if errArg1 != nil {
		panic(errArg1)
	}
	if errArg2 != nil {
		panic(errArg2)
	}
	if listSize < 1 || requestsNumber < 1 {
		fmt.Printf("Arguments must be numeric and positive.\n")
		return
	}

	utils.ExperimentRequests(listSize,requestsNumber)
	// utils.ExperimentListSizeEqualToRequestsWorstCase(listSize,requestsNumber)

}