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

	connections := utils.GraphAsSliceOfSlices(graph)
	for origin_it, conn := range connections {
		fmt.Printf("%d :", origin_it)
		fmt.Println(conn)
	}

	g := utils.DepthFirstSearch(nodes,connections)
	v := utils.TopologicalSorting(g) 
	fmt.Println()
	for _, verticinho := range v{
		fmt.Printf("\n\tColor = %s Parent = %d Discovered = %d Finished = %d Label = %s", string(verticinho.Color), verticinho.Parent, verticinho.Discovered, verticinho.Finished, verticinho.Label)
	}
}