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

func removeFromSlice(arrayWordCount []wordCount, itToRemove int) []wordCount {
	// Remove the element at index itToRemove from the slice
	previous := arrayWordCount[:itToRemove]
	following := arrayWordCount[itToRemove+1:]
	final := append(previous, following...)
	return final
}

func mergeCounts(arr_1, arr_2 []wordCount) []wordCount {
	// For each wordCount in arr_1:
	for i, wC_1 := range arr_1 {
		var found int = -1
		// Check if the wordCount exists in arr_2
		for j, wC_2 := range arr_2 {
			if wC_1.word == wC_2.word {
				// Modify arr_1 directly
				arr_1[i].count += wC_2.count
				found = j
				break
			}
		}
		// If found, remove the element from arr_2
		if found > -1 {
			arr_2 = removeFromSlice(arr_2, found)
		}
	}
	// Add remaining elements of arr_2 to arr_1
	return append(arr_1, arr_2...)
}

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

func countHelper(text string) []wordCount {

    // Get Rid Of Punctuation
    textNoPunctuation := removePunctuationMarks(text)

    // Split Text By Space
    words := splitText(textNoPunctuation)

    // Count Words
    resultDict := countWords(words)

    // Sort Array Descending 
    sortDict(resultDict)

    return resultDict
}

func main() {
    var casmurroInput = "./samples/dom-casmurro.txt"
    var memoriasInput = "./samples/memorias-postumas.txt"

    // Load text file as string
    textCasmurro := loadText(casmurroInput)
    
    // Count words
    wordsCasmurro := countHelper(textCasmurro)

    // Load another text file as string
    textMemorias := loadText(memoriasInput)
    wordsMemoria := countHelper(textMemorias)
    
    // Merge
    words := mergeCounts(wordsCasmurro, wordsMemoria)

    // Sort
    sortDict(words)

    // Write to local File
    writeWordCountDict(words, "./results/machado-array.txt")
}