package main

import (
	"fmt"
	"strconv"

	"github.com/one-cold-canuck/advent2024/internal/advent"
)

type Formatter struct {
	lineLength  int
	separator   string
	outlineChar string
}

func main() {

	formatter := Formatter{
		80,
		"*------------------------------------------------------------------------------*",
		"|",
	}

	PrintHeader(formatter)
	resultOne := advent.RunDayOne()
	resultTwo := advent.RunDayTwo()
	resultThree := advent.RunDayThree()
	resultFour := advent.RunDayFour()

	PrintResult(resultOne, formatter, 1)
	PrintResult(resultTwo, formatter, 2)
	PrintResult(resultThree, formatter, 3)
	PrintResult(resultFour, formatter, 4)
}

func PrintHeader(format Formatter) {
	fmt.Println(format.separator)
	fmt.Printf("|                               Advent of Code 2024!                           |\n")
	fmt.Println(format.separator)
}

func PrintResult(answer []int, format Formatter, day int) {

	dayCell := format.outlineChar + " Day " + strconv.Itoa(day) + " "
	if day < 10 {
		dayCell += " "
	}
	dayCell += format.outlineChar

	lineRemaining := format.lineLength - len(dayCell) - 1

	partOneCell := " Part 1: " + strconv.Itoa(answer[0])
	partTwoCell := " Part 2: " + strconv.Itoa(answer[1])

	cellCloser := ""

	for i := 0; i < (lineRemaining/2)-len(partOneCell); i++ {
		cellCloser += " "
	}

	partOneCell = partOneCell + cellCloser + format.outlineChar

	cellCloser = ""

	for i := 0; i < (lineRemaining/2)-len(partTwoCell); i++ {
		cellCloser += " "
	}

	partTwoCell = partTwoCell + cellCloser + format.outlineChar

	fmt.Printf("%s%s%s\n", dayCell, partOneCell, partTwoCell)
	fmt.Println(format.separator)
}
