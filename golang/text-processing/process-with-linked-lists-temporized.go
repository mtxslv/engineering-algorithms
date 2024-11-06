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

func writeWordCountDict(wordCountDict *linkedlist.LinkedList[wordCount], outputPath string) {
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	current := wordCountDict.Head()
	for current != nil {
		text := fmt.Sprintf("%s %d\n", current.Value().word, current.Value().count)
		_, err = f.WriteString(text)
		if err != nil {
			log.Fatal(err)
		}
		current = current.Next()
	}
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
	writeWordCountDict(mergedWords, "./results/machado-linkedlist.txt")
}
