package advent

import (
	"github.com/one-cold-canuck/advent2024/internal/util"
)

var dirs [][]int = [][]int{
	{-1, -1}, // up-left
	{0, -1},  // up
	{1, -1},  // up-right
	{-1, 0},  // left
	{1, 0},   // right
	{-1, 1},  // down-left
	{0, 1},   // down
	{1, 1},   // down-right
}

var diagdirs [][]int = [][]int{
	{-1, -1},
	{1, -1},
	{-1, 1},
	{1, 1},
}

var words []string = []string{"X", "M", "A", "S"}
var diagwords []string = []string{"M", "A", "S"}

func RunDayFour() (results []uint64) {

	results = append(results, 0)
	results = append(results, 0)

	input := util.ReadFileIntoStrings("data/day4.txt", "\n")
	input = input[:len(input)-1]

	xmascount := 0

	for x, row := range input {
		for y, character := range row {
			if string(words[0]) == string(character) {
				for _, dir := range dirs {
					if GetMatching(x, y, 1, dir, input) {
						xmascount += 1
					}
				}
			}
		}
	}

	results[0] = uint64(xmascount)

	diagxmascount := 0

	for x, row := range input {
		for y, character := range row {
			if string(character) == "A" {
				if GetMatchingDiag(x, y, input) {
					diagxmascount += 1
				}
			}
		}
	}
	results[1] = uint64(diagxmascount)

	return results
}

func GetMatching(xPos int, yPos int, wordPos int, dir []int, input []string) bool {
	nextX := xPos + dir[0]
	nextY := yPos + dir[1]

	if wordPos > len(words)-1 {
		return true
	}

	if nextX < 0 || nextX >= len(input) || nextY < 0 || nextY >= len(input[xPos]) {
		return false
	}

	if string(input[nextX][nextY]) == words[wordPos] {
		return GetMatching(nextX, nextY, wordPos+1, dir, input)
	}
	return false
}

func GetMatchingDiag(xPos int, yPos int, input []string) bool {

	if xPos-1 < 0 || xPos+1 >= len(input) || yPos-1 < 0 || yPos+1 >= len(input[xPos]) {
		return false
	}

	firstMatch := (string(input[xPos-1][yPos-1]) == "M" && string(input[xPos+1][yPos+1]) == "S") || (string(input[xPos-1][yPos-1]) == "S" && string(input[xPos+1][yPos+1]) == "M")
	secondMatch := (string(input[xPos-1][yPos+1]) == "M" && string(input[xPos+1][yPos-1]) == "S") || (string(input[xPos-1][yPos+1]) == "S" && string(input[xPos+1][yPos-1]) == "M")

	return firstMatch && secondMatch
}
