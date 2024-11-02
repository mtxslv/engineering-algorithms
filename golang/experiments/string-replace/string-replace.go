package main

import (
	"strings"
	"fmt"
)

func removePunctuationMarks(text string) string {
    punctuationMarks := []string{".", ",", "?", "!", "...", "_", "-", ":", ";", "\n", "\t", "--", "«", "»"}
    processedString := text
    for _, mark := range punctuationMarks {
        processedString = strings.ReplaceAll(processedString, mark, " ")
    }
    return processedString
}

func main() {
    text := "Hello, world! This is a test...\nLet's remove punctuation: shall we?"
    fmt.Println(removePunctuationMarks(text))
}