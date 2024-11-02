package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
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

func removePunctuationMarks(text string) string {
    punctuationMarks := []string{".", ",", "?", "!", "...", "_", "-", ":", ";", "\n", "\t", "--", "«", "»"}
    processedString := text
    for _, mark := range punctuationMarks {
        processedString = strings.ReplaceAll(processedString, mark, " ")
    }
    return processedString
}

func writeStringToFile(text string) {
    f, err := os.Create("./debug.txt")
    check(err)
    defer f.Close()
    _, err = f.WriteString(text)
    check(err)
    f.Sync()
}

func main() {
    var casmurroPath = "./samples/dom-casmurro.txt"
	domCasmurroText := loadText(casmurroPath)
    domCasmurroNoPunctuation := removePunctuationMarks(domCasmurroText)
    fmt.Printf("Found %d characters in processed text.\n", len(domCasmurroNoPunctuation))
    writeStringToFile(domCasmurroNoPunctuation)

    // var memoriasPath = "./samples/memorias-postumas.txt"
	// memoriasPostumasText := loadText(memoriasPath)
	// fmt.Print(memoriasPostumasText)
}