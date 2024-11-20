package utils

import (
	"io"
	"log"
	"os"
    "strings"
)


func LoadText(textPath string) string {
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

func BreakTextInNewLines(text string) []string {
    splittedString := strings.Split(text,"\n")
    var lines []string
    for _, splitted := range splittedString {
        // Remove four-whitespaces prefix
        splitted = strings.TrimPrefix(splitted, "    ")
        // Remove empty characters
        if len(splitted) > 0 {
            lines = append(lines, splitted)
        }
    }
    return lines
}