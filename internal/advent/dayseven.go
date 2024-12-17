package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/one-cold-canuck/advent2024/internal/util"
)

type calibration struct {
	Result   int
	Operands []int
	IsValid  bool
}

var calibrations []calibration

func RunDaySeven() []int {

	results := []int{0, 0}

	puzzleInput := util.ReadFileIntoStrings("data/day7.txt", "\n")
	puzzleInput = puzzleInput[:len(puzzleInput)-1]

	for _, v := range puzzleInput {
		calibrations = append(calibrations, ProcessCalibration(v))
	}

	for _, v := range calibrations {

		CheckCalibration(v)
	}
	return results

}

func ProcessCalibration(input string) calibration {

	var caldata calibration

	elements := strings.Split(input, " ")

	caldata.Result, _ = strconv.Atoi(strings.TrimRight(elements[0], ":"))
	caldata.IsValid = false

	for i, v := range elements {
		if i != 0 {
			intvalue, _ := strconv.Atoi(v)
			caldata.Operands = append(caldata.Operands, intvalue)
		}
	}

	return caldata

}

func CheckCalibration(input calibration) {
	operatorCount := len(input.Operands) - 1
	fmt.Println("Calibration: ", input.Result, ", NumberofElements: ", operatorCount, ", Factorial: ", util.Factorial(operatorCount))
}
