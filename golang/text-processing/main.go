package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// WORD COUNT STRUCT

type wordCount struct {
    word string 
    count uint64 
}

// SAVE FILE

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func writeStringToFile(text string) {
    // modified from https://gobyexample.com/writing-files
    f, err := os.Create("./debug.txt")
    check(err)
    defer f.Close()
    _, err = f.WriteString(text)
    check(err)
    f.Sync()
}

// READ FILE

func loadText(textPath string) string {
	// adapted from https://stackoverflow.com/questions/36111777/how-to-read-a-text-file
	// and from https://stackoverflow.com/questions/9644139/from-io-reader-to-string-in-go
    file, err := os.Open(textPath)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()
  b, err := io.ReadAll(file)
  text := string(b)
  return text
}

// TEXT PROCESSING

func removePunctuationMarks(text string) string {
    punctuationMarks := []string{".", ",", "?", "!", "...", "_", "-", ":", ";", "\n", "\r", "\t","\ufeff", "--", "«", "»"}
    processedString := text
    for _, mark := range punctuationMarks {
        processedString = strings.ReplaceAll(processedString, mark, " ")
    }
    return processedString
}

func splitText(text string) []string {
    splittedString := strings.Split(text," ")
    var words []string
    for _, splitted := range splittedString {
        // Remove empty characters
        if len(splitted) > 0 {
            words = append(words, splitted)
        }
    }
    return words
}


// MAIN FUNCTION 

func main() {
    var casmurroPath = "./samples/dom-casmurro.txt"
	domCasmurroText := loadText(casmurroPath)
    domCasmurroNoPunctuation := removePunctuationMarks(domCasmurroText)
    fmt.Printf("Found %d characters in processed text.\n", len(domCasmurroNoPunctuation))
    words := splitText(domCasmurroNoPunctuation)
    fmt.Printf("THERE ARE %d WORDS IN TOTAL \n", len(words))
    fmt.Printf("SAMPLE:%q \n", words[:100])
}