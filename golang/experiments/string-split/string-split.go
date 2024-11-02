package main

import (
	"fmt"
	"strings"
)

func main(){
	var text string = "o rato roeu a roupa do rei de roma"
	words := strings.Split(text," ")
	fmt.Printf("%q\n\n",words)
	for i, word := range words {
		fmt.Printf("%d-th Word: %s (%d characters)\n",i, word, len(word))
	}
}
