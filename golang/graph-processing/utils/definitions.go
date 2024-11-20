package utils

type nodeNameAndLabel struct {
    nodeName string
    nodeLabel string
} 

type graphEdge struct {
    nodeNameOrigin string
    nodeNameDestiny string
}

type simpleGraph struct {
	Nodes []nodeNameAndLabel
	Edges []graphEdge
}