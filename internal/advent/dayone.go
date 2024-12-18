package advent

import (
	"sort"

	"github.com/one-cold-canuck/advent2024/internal/util"
)

// RunDayOne processes the input for the Day 1 AoC puzzle (https://adventofcode.com/2024/day/1)
// Input is split into 2 lists.
func RunDayOne() []uint64 {

	results := []uint64{}

	rightList := util.ReadFileIntoInts("data/day1_rightlist.txt", "\n")
	leftList := util.ReadFileIntoInts("data/day1_leftlist.txt", "\n")

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})
	count := 0
	matchCount := 0

	for i, v := range leftList {
		internalMatchCount := 0

		if rightList[i] > v {
			count += (rightList[i] - v)
		} else {
			count += (v - rightList[i])
		}
		for _, w := range rightList {
			if w == v {
				internalMatchCount += 1
			}
		}
		matchCount += v * internalMatchCount
	}

	results = append(results, uint64(count))
	results = append(results, uint64(matchCount))
	return results
}
