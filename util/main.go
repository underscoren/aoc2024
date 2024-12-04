package util

import (
	"io"
	"log"
	"os"
	"strings"
)

// Integer absolute
func Abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

// Reads the entire contents of a file as a string
func ReadFile(filepath string) string {
	// reading in the file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

// Parses a string of text into a slice of strings split by newlines
func GetLines(text string) []string {
	return strings.Split(strings.Trim(text, "\n"), "\n")
}

// Reads in a file as a slice of strings
func ReadLines(filepath string) []string {
	return GetLines(ReadFile(filepath))
}

// removes an element from a slice at the given position
// without modifying the original
func Remove(slice []int, position int) []int {
	copy := make([]int, 0, len(slice)-1)
	copy = append(copy, slice[:position]...)
	return append(copy, slice[position+1:]...)
}
