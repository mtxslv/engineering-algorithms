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
    count int 
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

func writeWordCountDict(wordCountDict []wordCount) {
    f, err := os.Create("./debug.txt")
    check(err)
    defer f.Close()
    for _, wordWithCounting := range wordCountDict{
        text := fmt.Sprintf("%s %d\n", wordWithCounting.word, wordWithCounting.count)
        _, err = f.WriteString(text)
    } 
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

// TEXT PRE-PROCESSING

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

// WORD COUNT

func findWordPosition(dictionary []wordCount, word string) int {
    for it, wordWithCount := range dictionary {
        if wordWithCount.word == word {
            return it
        }
    } 
    return -1
}
// 

func countWords(words []string) []wordCount {
    var wordCountDict []wordCount
    for _, word := range words {
        // First word?
        if len(wordCountDict) == 0 {
            first := wordCount{word, 1}
            wordCountDict = append(wordCountDict, first)
        } else {
            // check if word exist in counting Words Slice.
            wordWithCountPosition := findWordPosition(wordCountDict, word)
            // if so, increment count number
            if wordWithCountPosition > 0{
                wordCountDict[wordWithCountPosition].count++
            } else { // if not, append
                wordCountDict = append(wordCountDict, wordCount{word, 1})
            }        
            // insertion sort based on word key to keep search working
        }
    }
    return wordCountDict
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
    resultDict := countWords(words[:500])
    fmt.Printf("DICT:%q \n", resultDict)
    writeWordCountDict(resultDict)
}