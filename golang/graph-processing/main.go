package main

import (
	"fmt"
	"utils/utils"
)


func main(){
	var textPath = "graph.dot"
	song := utils.LoadText(textPath)
	fmt.Print(song)
	fmt.Printf("music to watch boys to \n")
}