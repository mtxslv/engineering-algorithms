package utils

import (
	"fmt"
	"sort"
)

type dfsVertex struct {
	Color rune // 'b' for black, 'g' for grey, and 'w' for white 
	Parent int // the parent's Nodes array position 
	Discovered int 
	Finished int 
	Label string // key to map to nodeNameAndLabel
}

type dfsGraph struct {
	Vertices []dfsVertex
	Connections [][]int
}

func SetDfsGraph(nodes []nodeNameAndLabel, connections [][]int) dfsGraph {
	var vertices []dfsVertex
	for _, node := range nodes {
		aVertex := dfsVertex{
			Color : 'w',
			Parent : -1,
			Discovered: -1,
			Finished : -1,
			Label : node.nodeLabel,
		}
		vertices = append(vertices, aVertex)
	} 
	return dfsGraph{Vertices: vertices, Connections: connections}
}

func DepthFirstSearch(nodes []nodeNameAndLabel, connections [][]int) dfsGraph {
    // Set Vertices' color, parent, F (and label) on Graph 
    graph := SetDfsGraph(nodes, connections)
    
    // Now follow DFS procedure
    var time = 0
    for vertex_it, vertex := range graph.Vertices {
        if vertex.Color == 'w' {
            // fmt.Print("VC EH BRANCO\n")
            DepthFirstSearchVisit(&graph, vertex_it, &time) // Pass a pointer to the graph
            // break
        }
		fmt.Printf("\n\tColor = %s Parent = %d Discovered = %d Finished = %d Label = %s", string(graph.Vertices[vertex_it].Color), graph.Vertices[vertex_it].Parent, graph.Vertices[vertex_it].Discovered, graph.Vertices[vertex_it].Finished, graph.Vertices[vertex_it].Label)
    }
    return graph 
}

func DepthFirstSearchVisit(graph *dfsGraph, vertex_it int, time *int) {
    // Increase time
    *time++
    graph.Vertices[vertex_it].Discovered = *time
    graph.Vertices[vertex_it].Color = 'g'

	for _, adj_vertex_id := range graph.Connections[vertex_it]{
		if graph.Vertices[adj_vertex_id].Color == 'w'{
			graph.Vertices[adj_vertex_id].Parent = vertex_it
			DepthFirstSearchVisit(graph, adj_vertex_id, time)
		}
	}

    *time++
    graph.Vertices[vertex_it].Finished = *time
    graph.Vertices[vertex_it].Color = 'b'
}

func TopologicalSorting(graph dfsGraph) []dfsVertex {
	v := make([]dfsVertex, len(graph.Vertices))
	copy(v, graph.Vertices)
	sort.Slice(v, func(i,j int) bool {
		return v[i].Finished > v[j].Finished
	})
	return v
}