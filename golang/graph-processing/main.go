package main

import (
	"fmt"
	"utils/utils"
)



func main(){
	var graphPath = "graph.dot"
	graphDefinition := utils.LoadGraphDefinition(graphPath)
	graph := utils.ExtractGraph(graphDefinition)

	nodes := graph.Nodes
	edges := graph.Edges

	for i, node := range nodes {
		fmt.Printf("NODE #%d: %+v\n", i+1, node)
	} 

	for i, edge := range edges {
		fmt.Printf("EDGE #%d: %+v\n", i+1, edge)
	} 	

	a := utils.GraphAsSliceOfSlices(graph)
	for origin_it, ass := range a {
		fmt.Printf("%d :", origin_it)
		fmt.Println(ass)
	}
	// fmt.Println(a)
}