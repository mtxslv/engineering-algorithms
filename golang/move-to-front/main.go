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
	if len(argsWithProg) != 4 {
		fmt.Printf("Usage: \n./main <type> <list-size> <number-of-requests>\n")
		fmt.Printf("\t <type> must be 'equal' or 'diff'\n")
		return
	} 
	// What is the type?
	typeArg := argsWithProg[1]

	if typeArg != "equal" && typeArg != "diff" {
		fmt.Printf("Usage: \n./main <type> <list-size> <number-of-requests>\n")
		fmt.Printf("\t <type> must be 'equal' or 'diff'\n")
		return		
	}
	// If everything is fine, code follows...

	// Define parameters
	listSize, errArg1 := strconv.Atoi(argsWithProg[2])
	requestsNumber, errArg2 := strconv.Atoi(argsWithProg[3])
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
	
	if typeArg == "diff" {
		utils.ExperimentRequests(listSize,requestsNumber)
	} else {
		utils.ExperimentListSizeEqualToRequestsWorstCase(listSize,requestsNumber)
	}

}