package main

import (
	"fmt"
)

func main(){
	var words []string = []string{"ave", "avestruz", "avestan", "avião", "aviação"}
	for it := 0 ; it < len(words)-1 ; it++ {
		current := words[it]
		next := words[it+1]
		fmt.Printf("%s <= %s?", current, next)
		if current <= next {
			fmt.Printf(" (Order Maintained)\n")
		} else {
			fmt.Printf(" (Order Broke)\n")
		}
	}
}
