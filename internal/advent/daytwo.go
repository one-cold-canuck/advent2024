package advent

import (
	"github.com/one-cold-canuck/advent2024/internal/util"
	"strconv"
	"strings"
)

func RunDayTwo() []int {

	results := []int{}

	data := util.ReadFileIntoStrings("data/day2.txt")

	for _, v := range data {
		dataline := strings.Split(v, " ")
		ascending := false
		descending := false
		changeDirection := false
		for i, val := range dataline {
			if i >= len(dataline)-1 {
				break
			}
			if val < dataline[i] {

			}
		}

	}
	results = append(results, 0)
	results = append(results, 0)

	return results
}
