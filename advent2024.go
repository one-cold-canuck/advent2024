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
	PrintResult(resultOne, formatter, 1)
	resultTwo := advent.RunDayTwo()
	PrintResult(resultTwo, formatter, 2)
	resultThree := advent.RunDayThree()
	PrintResult(resultThree, formatter, 3)
	resultFour := advent.RunDayFour()
	PrintResult(resultFour, formatter, 4)
	resultFive := advent.RunDayFive()
	PrintResult(resultFive, formatter, 5)
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
