package advent

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/one-cold-canuck/advent2024/internal/util"
)

func RunDayThree() (results []uint64) {

	input := util.ReadFile("data/day3.txt")

	re := regexp.MustCompile(`mul\([0-9]*,[0-9]*\)|do\(\)|don't\(\)`)
	validCommands := re.FindAll([]byte(input), -1)

	resultOne := 0
	resultTwo := 0

	multiply := true

	for _, v := range validCommands {

		value := string(v)
		if value == "do()" {
			multiply = true
		} else if value == "don't()" {
			multiply = false
		} else {
			value, _ = strings.CutPrefix(value, "mul(")
			value, _ = strings.CutSuffix(value, ")")

			arguments := strings.Split(value, ",")
			firstArg, _ := strconv.Atoi(arguments[0])
			secondArg, _ := strconv.Atoi(arguments[1])

			resultOne += firstArg * secondArg
			if multiply {
				resultTwo += firstArg * secondArg
			}
		}
	}

	results = append(results, uint64(resultOne))
	results = append(results, uint64(resultTwo))

	return results
}
