package main

import (
	"io"
	"fmt"
	"log"
	"os"
	"time"
	"strings"
	"text-processing/linkedlist" // Assuming your linked list package is imported here
)

type wordCount struct {
	word  string
	count int
}

// UTILS

func wordCountPartition(A []wordCount, p int, r int) int {
	x := A[r].count // the pivot
	i := p - 1 // highest index into the low side
	for j := p ; j < r ; j++ {
		if A[j].count > x { // sort descending
			i = i + 1 
			aux := A[i]
			A[i] = A[j]
			A[j] = aux
		}
	} 
	aux := A[r]
	A[r] = A[i+1]
	A[i+1] = aux
	return i + 1
}

func wordCountQuicksort(A []wordCount, p int, r int){
	if p < r {
		// Partition the subarray around the pivot, which ends up in A[q].
		q := wordCountPartition(A,p,r)
		wordCountQuicksort(A, p, q-1) // recursively sort the low side
		wordCountQuicksort(A, q+1, r) // recursively sort the high side
	}
}

func sortDict(dict []wordCount) {
    wordCountQuicksort(dict,0,len(dict)-1)
}


func findWordPosition(ll *linkedlist.LinkedList[wordCount], word string) (*linkedlist.Node[wordCount], int64) {
    var totalTime int64 
    timeStart := time.Now()	
	current := ll.Head()
	for current != nil {
		if current.Value().word == word {
			totalTime = time.Since(timeStart).Nanoseconds()
			return current, totalTime
		}
		current = current.Next()
	}
    totalTime = time.Since(timeStart).Nanoseconds()
	return nil, totalTime
}

func mergeCounts(ll1, ll2 *linkedlist.LinkedList[wordCount]) *linkedlist.LinkedList[wordCount] {
	// For each wordCount in ll2, merge counts into ll1
	current_ll2 := ll2.Head()
	var totalRemovalTime int64 = 0

	for current_ll2 != nil {
		node, _ := findWordPosition(ll1, current_ll2.Value().word)
		if node != nil {
			node.Value().count += current_ll2.Value().count
			//  remove node from ll1
			startRemove := time.Now()
			ll1.Remove(*node.Value())
			timeToRemove := time.Since(startRemove)
            totalRemovalTime += timeToRemove.Nanoseconds()
		} else {
			ll1.Add(*current_ll2.Value())
		}
		current_ll2 = current_ll2.Next()
	}
	fmt.Printf("Total Removal Time (ns): %d\n", totalRemovalTime)
	return ll1
}
// TEXT PRE-PROCESSING

func removePunctuationMarks(text string) string {
	punctuationMarks := []string{".", ",", "?", "!", "...", "_", "-", ":", ";", "\n", "\r", "\t", "\ufeff", "--", "«", "»"}
	processedString := text
	for _, mark := range punctuationMarks {
		processedString = strings.ReplaceAll(processedString, mark, " ")
	}
	return processedString
}

func splitText(text string) []string {
	splittedString := strings.Split(text, " ")
	var words []string
	for _, splitted := range splittedString {
		if len(splitted) > 0 {
			words = append(words, splitted)
		}
	}
	return words
}

// WORD COUNT

func countWords(words []string) *linkedlist.LinkedList[wordCount] {
	wordCountDict := linkedlist.New[wordCount]()
    var totalSearchTime int64 = 0
	for _, word := range words {
		node, searchTime := findWordPosition(wordCountDict, word)
		totalSearchTime += searchTime
		if node != nil {
			node.Value().count++
		} else {
			wordCountDict.Add(wordCount{word, 1})
		}
	}
    fmt.Printf("Total Search Time: %d \n", totalSearchTime)
	return wordCountDict
}

// SAVE FILE

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func writeWordCountDict(wordCountDict []wordCount, outputPath string) {
    f, err := os.Create(outputPath)
    check(err)
    defer f.Close()
    for _, wordWithCounting := range wordCountDict{
        text := fmt.Sprintf("%s %d\n", wordWithCounting.word, wordWithCounting.count)
        _, err = f.WriteString(text)
    } 
    check(err)
    f.Sync()    
}

// MAIN FUNCTION

func countHelper(text string) *linkedlist.LinkedList[wordCount] {
	textNoPunctuation := removePunctuationMarks(text)
	words := splitText(textNoPunctuation)

	// Count Words
    start := time.Now()
    resultDict := countWords(words)
    countWordsTimeElapsed := time.Since(start).Nanoseconds()
    fmt.Printf("TIME ELAPSED (ns): %d\n", countWordsTimeElapsed)

	return resultDict
}

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

func main() {
	casmurroInput := "./samples/dom-casmurro.txt"
	memoriasInput := "./samples/memorias-postumas.txt"

	fmt.Printf("\t PROCESSANDO CASMURRO \n")

	textCasmurro := loadText(casmurroInput)
	wordsCasmurro := countHelper(textCasmurro)

    fmt.Printf("\t PROCESSANDO MEMORIAS \n")
	textMemorias := loadText(memoriasInput)
	wordsMemorias := countHelper(textMemorias)
	
    fmt.Printf("\t MERGE \n")
	mergedWords := mergeCounts(wordsCasmurro, wordsMemorias)

	words := mergedWords.ToArray()

	sortDict(words)

	writeWordCountDict(words, "./results/machado-linkedlist.txt")
}
