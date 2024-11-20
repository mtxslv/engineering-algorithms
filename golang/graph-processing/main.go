package main

import (
	"fmt"
	"utils/utils"
)



func main(){
	var graphPath = "graph.dot"
	graphDefinition := utils.LoadGraphDefinition(graphPath)
	nodes, edges := utils.ExtractNodesAndEdges(graphDefinition)

	for i, node := range nodes {
		fmt.Printf("NODE #%d: %+v\n", i+1, node)
	} 

	for i, edge := range edges {
		fmt.Printf("EDGE #%d: %+v\n", i+1, edge)
	} 	
}