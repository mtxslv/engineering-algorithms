package main

import (
	"io"
	"fmt"
	"log"
	"os"
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


func findWordPosition(ll *linkedlist.LinkedList[wordCount], word string) *linkedlist.Node[wordCount] {
	current := ll.Head()
	for current != nil {
		if current.Value().word == word {
			return current
		}
		current = current.Next()
	}
	return nil
}

func mergeCounts(ll1, ll2 *linkedlist.LinkedList[wordCount]) *linkedlist.LinkedList[wordCount] {
	// For each wordCount in ll2, merge counts into ll1
	current := ll2.Head()
	for current != nil {
		node := findWordPosition(ll1, current.Value().word)
		if node != nil {
			node.Value().count += current.Value().count
		} else {
			ll1.Add(*current.Value())
		}
		current = current.Next()
	}
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
	for _, word := range words {
		node := findWordPosition(wordCountDict, word)
		if node != nil {
			node.Value().count++
		} else {
			wordCountDict.Add(wordCount{word, 1})
		}
	}
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
	return countWords(words)
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

	textCasmurro := loadText(casmurroInput)
	wordsCasmurro := countHelper(textCasmurro)

	textMemorias := loadText(memoriasInput)
	wordsMemorias := countHelper(textMemorias)

	mergedWords := mergeCounts(wordsCasmurro, wordsMemorias)

	words := mergedWords.ToArray()

	sortDict(words)

	writeWordCountDict(words, "./results/machado-linkedlist.txt")
}
