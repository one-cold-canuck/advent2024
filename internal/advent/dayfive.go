package advent

import (
	"strconv"
	"strings"

	"github.com/one-cold-canuck/advent2024/internal/util"
)

var invalidprints []string

func RunDayFive() []uint64 {

	results := []uint64{0, 0}

	ruleset := GetRules("data/day5-order.txt")
	printset := GetPrintset("data/day5-updates.txt")

	validprintCount := 0
	invalidprintCount := 0

	for _, print := range printset {
		if CheckPrintset(print, ruleset) {
			pages := strings.Split(print, ",")
			middleIndex := len(pages) / 2
			middleValue, _ := strconv.Atoi(pages[middleIndex])
			validprintCount += middleValue
		} else {
			invalidprints = append(invalidprints, print)
		}
	}

	for _, print := range invalidprints {
		sortedlist, iscorrect := SortPrintset(print, ruleset)
		for !iscorrect {
			sortedlist, iscorrect = SortPrintset(sortedlist, ruleset)
		}
		pages := strings.Split(sortedlist, ",")
		middleIndex := len(pages) / 2
		middleValue, _ := strconv.Atoi(pages[middleIndex])
		invalidprintCount += middleValue
	}

	results[0] = uint64(validprintCount)
	results[1] = uint64(invalidprintCount)

	return results
}

func GetRules(filename string) [][]string {

	input := util.ReadFileIntoStrings(filename, "\n")
	input = input[:len(input)-1]

	ruleset := make([][]string, 0)

	for _, v := range input {
		rule := strings.Split(v, "|")
		ruleset = append(ruleset, rule)
	}

	return ruleset
}

func GetPrintset(filename string) []string {
	printset := util.ReadFileIntoStrings("data/day5-updates.txt", "\n")
	printset = printset[:len(printset)-1]

	return printset
}

func CheckPrintset(printlist string, ruleset [][]string) bool {
	for _, rule := range ruleset {
		beforepageIndex := strings.Index(printlist, rule[0])
		afterpageIndex := strings.Index(printlist, rule[1])
		if afterpageIndex >= 0 && beforepageIndex >= 0 {
			if afterpageIndex < beforepageIndex {
				return false
			}
		}
	}
	return true
}

func SortPrintset(printlist string, ruleset [][]string) (string, bool) {

	for _, rule := range ruleset {
		beforepageIndex := strings.Index(printlist, rule[0])
		afterpageIndex := strings.Index(printlist, rule[1])
		if afterpageIndex >= 0 && beforepageIndex >= 0 {
			if afterpageIndex < beforepageIndex {
				r := strings.NewReplacer(rule[0], rule[1], rule[1], rule[0])
				printlist = r.Replace(printlist)
				return printlist, false
			}
		}
	}
	return printlist, true
}
