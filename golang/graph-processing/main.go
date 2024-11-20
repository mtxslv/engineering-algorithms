package main

import (
	"fmt"
	"utils/utils"
)



func main(){
	var graphPath = "graph.dot"
	graphDefinition := utils.LoadGraphDefinition(graphPath)
	for i, cmd := range graphDefinition {
		fmt.Printf("%d LINE: \n", i)
		fmt.Println(cmd)
	} 
}