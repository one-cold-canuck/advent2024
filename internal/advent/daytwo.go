package advent

import (
	"slices"
	"strconv"

	"strings"

	"github.com/one-cold-canuck/advent2024/internal/util"
)

func RunDayTwo() []uint64 {

	results := []uint64{}

	data := util.ReadFileIntoStrings("data/day2.txt", "\n")
	counter := 0  // for Part 1 output
	counter2 := 0 // For Part 2 output
	data = data[:len(data)-1]

	for _, v := range data {

		v = strings.Trim(v, "\r\n")
		datalineS := strings.Split(v, " ")
		dataline := make([]int, 0)

		for i := range datalineS {
			intval, _ := strconv.Atoi(datalineS[i])
			dataline = append(dataline, intval)
		}

		if CheckValid(dataline) {
			counter += 1
			counter2 += 1
		} else {
			for i := range dataline {
				dtl := make([]int, 0)
				for i := range dataline {
					dtl = append(dtl, dataline[i])
				}
				dtl = slices.Delete(dtl, i, i+1)
				if CheckValid(dtl) {
					counter2 += 1
					break
				}

			}
		}

	}
	results = append(results, uint64(counter))
	results = append(results, uint64(counter2))

	return results
}

func CheckValid(dataline []int) bool {

	isAscending := true
	isDescending := true
	tooDistant := false

	lastElement := len(dataline) - 2

	for i := 0; i <= lastElement; i++ {
		if !(dataline[i] < dataline[i+1]) && isAscending {
			isAscending = false
		}

		if !(dataline[i] > dataline[i+1]) && isDescending {
			isDescending = false
		}
		if dataline[i]-dataline[i+1] > 3 || dataline[i+1]-dataline[i] > 3 {
			tooDistant = true
		}
	}

	dataIsGood := isAscending != isDescending

	return dataIsGood && !tooDistant
}
