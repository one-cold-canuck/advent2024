// Package util provides utility functions for Advent of Code
package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads a file of newline-separated integers, and returns a Slice of Ints
func ReadFile(filePath string) (numbers []int) {

	numbers = []int{}

	values, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file!")
		return
	}

	lines := strings.Split(string(values), "\n")

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
