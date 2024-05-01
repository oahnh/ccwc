package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {

	var flagvar string
	flag.StringVar(&flagvar, "c", "", "Counts number of bytes in the file")
	flag.StringVar(&flagvar, "l", "", "Counts number of lines in the file")
	flag.StringVar(&flagvar, "w", "", "Counts number of words in the file")
	flag.StringVar(&flagvar, "m", "", "Counts number of characters in the file")

	flag.Parse()

	fileName, data := parseFile()

	// to identify flag, use in switch case
	var flagName string
	flag.Visit(func(f *flag.Flag) {
		flagName = f.Name
	})

	if flagName == "c" {
		byteCount := getByteCount(data)
		fmt.Printf("%d %s\n", byteCount, fileName)
	} else if flagName == "l" {
		lineCount := getLineCount(data)
		fmt.Printf("%d %s\n", lineCount, fileName)
	} else if flagName == "w" {
		wordCount := getWordCount(data)
		fmt.Printf("%d %s\n", wordCount, fileName)
	} else if flagName == "m" {
		charCount := getChars(data)
		fmt.Printf("%d %s\n", charCount, fileName)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getByteCount(b []byte) int {
	return len(b)
}

func getLineCount(b []byte) int {
	return bytes.Count(b, []byte{'\n'})
}

func getWordCount(b []byte) int {
	return len(bytes.Fields(b))
}

func getChars(b []byte) int {
	return utf8.RuneCountInString(string(b))
}

func parseFile() (string, []byte) {
	if len(os.Args) < 2 {
		panic("Expected file path to process")
	}

	name := os.Args[2]
	data, err := os.ReadFile(name)
	check(err)

	return name, data
}
