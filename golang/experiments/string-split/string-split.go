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
	fmt.Println("--------------------------------------")
	// Now multi line string
	var stanza string = `Me dá um tempo pra pensar
						 Tenho que criar coragem pra me arriscar
						 Por Favor entenda o que eu 'to passando
						 Só 'to pedindo um tempo, não 'to recusando`
	filteredStanza := strings.Replace(stanza, "\t","",-1)
	verses := strings.Split(filteredStanza, "\n")
	fmt.Printf("%q\n\n",verses)
	var otherWords []string
	for i, verse := range verses {
		fmt.Printf("Processing verse #%d\n", i)
		otherWords = append(otherWords, strings.Split(verse, " ")...)
	}
	fmt.Printf("%q\n\n",otherWords)

}
