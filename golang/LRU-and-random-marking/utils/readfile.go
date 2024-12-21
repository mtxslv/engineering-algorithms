package utils

import (
	"io"
	"log"
	"os"
	"strings"
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

func LoadLines(textpath string) []string {
	// read full text as string
	text := loadText(textpath)

	// split text by lines
	spplited := strings.Split(text,"\n")
	var words []string
	
	// remove words whose size is 0
	for _, str := range spplited {
		if len(str) > 0 {
			words = append(words, str)
		}
	} 
	return words
}