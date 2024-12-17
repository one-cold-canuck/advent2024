// Package util provides utility functions for Advent of Code
package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFileIntoInts reads a file of newline-separated integers, and returns a Slice of Ints
func ReadFileIntoInts(filePath string, separator string) (numbers []int) {

	numbers = []int{}

	values, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file!")
		return
	}

	lines := strings.Split(string(values), separator)

	for i, v := range lines {
		if i < len(lines)-1 {
			value, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Error parsing file at index %d!\n", i)
				return
			}

			numbers = append(numbers, value)
		}
	}

	return numbers
}

func ReadFileIntoStrings(filePath string, separator string) (results []string) {
	results = []string{}

	values, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file!")
		return
	}

	lines := strings.Split(string(values), separator)

	return lines
}

func ReadFile(filePath string) (filecontent []byte) {

	filecontent, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file!")
		return nil
	}

	return filecontent
}
