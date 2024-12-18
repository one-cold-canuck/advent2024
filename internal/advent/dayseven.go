package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/one-cold-canuck/advent2024/internal/util"
)

type calibration struct {
	Result   uint64
	Operands []uint64
	IsValid  bool
}

var calibrations []calibration

func RunDaySeven() []uint64 {

	results := []uint64{0, 0}

	puzzleInput := util.ReadFileIntoStrings("data/day7.txt", "\n")
	puzzleInput = puzzleInput[:len(puzzleInput)-1]

	for _, v := range puzzleInput {
		calibrations = append(calibrations, ProcessCalibration(v))
	}

	puzzleInput = nil

	for i, _ := range calibrations {
		CheckCalibration(i)
		if calibrations[i].IsValid {
			results[0] += uint64(calibrations[i].Result)
		}
	}

	for j, _ := range calibrations {
		calibrations[j].IsValid = false
	}

	for k, _ := range calibrations {
		CheckCalibrationThreeOp(k)
		if calibrations[k].IsValid {
			results[1] += uint64(calibrations[k].Result)
		}
	}
	return results

}

func ProcessCalibration(input string) calibration {

	var caldata calibration

	elements := strings.Split(input, " ")
	caldata.Result, _ = strconv.ParseUint(strings.TrimRight(elements[0], ":"), 10, 64)
	caldata.IsValid = false

	for i, v := range elements {
		if i != 0 {
			intvalue, _ := strconv.Atoi(v)
			caldata.Operands = append(caldata.Operands, uint64(intvalue))
		}
	}

	return caldata

}
func CheckCalibration(index int) {
	operatorCount := uint64(len(calibrations[index].Operands) - 1)
	opbits := ""
	for i := operatorCount; i > 0; i-- {
		opbits += "1"
	}
	strconv.ParseInt(opbits, 2, 64)
	for j, _ := strconv.ParseInt(opbits, 2, 64); j >= 0; j-- {
		currbits := strconv.FormatInt(j, 2)
		for len(currbits) < int(operatorCount) {
			currbits = "0" + currbits
		}
		checkValue := calibrations[index].Operands[0]
		for opindex, v := range currbits {
			switch v {
			case '1':
				checkValue = checkValue * calibrations[index].Operands[opindex+1]
			case '0':
				checkValue = checkValue + calibrations[index].Operands[opindex+1]
			}
		}
		if checkValue == calibrations[index].Result {
			calibrations[index].IsValid = true
			return
		}
	}
}

func CheckCalibrationThreeOp(index int) {
	operatorCount := int64(len(calibrations[index].Operands) - 1)
	opbits := ""
	for i := operatorCount * 2; i > 0; i-- {
		opbits += "1"
	}
	fmt.Println("Opbits: ", opbits)
	strconv.ParseInt(opbits, 2, 64)
	for j, _ := strconv.ParseInt(opbits, 2, 64); j >= 0; j-- {
		currbits := strconv.FormatInt(j, 2)
		for len(currbits) < int(operatorCount*2) {
			currbits = "0" + currbits
		}
		checkValue := calibrations[index].Operands[0]
		for opindex := 0; opindex < len(currbits); opindex += 2 {
			op := currbits[opindex : opindex+2]
			switch op {
			case "10":
				checkValue = checkValue + calibrations[index].Operands[(opindex/2)+1]
			case "01":
				checkValue = checkValue * calibrations[index].Operands[(opindex/2)+1]
			case "00":
				startval := strconv.FormatUint(checkValue, 10)
				nextval := strconv.FormatUint(calibrations[index].Operands[(opindex/2)+1], 10)
				newval := startval + nextval
				checkValue, _ = strconv.ParseUint(newval, 10, 64)
			}
		}
		if checkValue == calibrations[index].Result {
			fmt.Println("Found Match: ", calibrations[index].Result, ", Operands: ", calibrations[index].Operands, "Calculation: ", checkValue)
			calibrations[index].IsValid = true
			return
		}
	}
}
