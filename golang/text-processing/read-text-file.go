package main

import (
    "fmt"
    "io"
    "os"
    "log"
)

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
	var memoriasPath = "./samples/memorias-postumas.txt"
	var casmurroPath = "./samples/dom-casmurro.txt"
	memoriasPostumasText := loadText(memoriasPath)
	domCasmurroText := loadText(casmurroPath)
	fmt.Print(memoriasPostumasText)
	fmt.Print(domCasmurroText)
}