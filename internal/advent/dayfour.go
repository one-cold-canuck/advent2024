package advent

import (
	"fmt"
	"regexp"
	"slices"

	"github.com/one-cold-canuck/advent2024/internal/util"
)

func RunDayFour() (results []int) {

	results = append(results, 0)
	results = append(results, 0)

	input := util.ReadFileIntoStrings("data/day4.txt")

	gameboard := make([][]rune, 0)

	for _, v := range input {
		gameline := make([]rune, 0)
		for _, val := range v {
			if val != '\n' {
				gameline = append(gameline, val)
			}
		}
		gameboard = append(gameboard, gameline)
	}

	if len(gameboard) > 0 {
		gameboard = gameboard[:len(gameboard)-1]
	}

	re := regexp.MustCompile("XMAS|SAMX")

	COL_BOUND := len(gameboard[0]) - 4
	ROW_BOUND := len(gameboard) - 4

	xmasCount := 0
	rowsChecked := make([]int, 0)

	for i := 0; i <= ROW_BOUND; i++ {
		for j := 0; j <= COL_BOUND; j++ {

			rowone := []rune{gameboard[i][j], gameboard[i][j+1], gameboard[i][j+2], gameboard[i][j+3]}
			rowtwo := []rune{gameboard[i+1][j], gameboard[i+1][j+1], gameboard[i+1][j+2], gameboard[i+1][j+3]}
			rowthree := []rune{gameboard[i+2][j], gameboard[i+2][j+1], gameboard[i+2][j+2], gameboard[i+2][j+3]}
			rowfour := []rune{gameboard[i+3][j], gameboard[i+3][j+1], gameboard[i+3][j+2], gameboard[i+3][j+3]}

			colone := []rune{gameboard[i][j], gameboard[i+1][j], gameboard[i+2][j], gameboard[i+3][j]}
			coltwo := []rune{gameboard[i][j+1], gameboard[i+1][j+1], gameboard[i+2][j+1], gameboard[i+3][j+1]}
			colthree := []rune{gameboard[i][j+2], gameboard[i+1][j+2], gameboard[i+2][j+2], gameboard[i+3][j+2]}
			colfour := []rune{gameboard[i][j+3], gameboard[i+1][j+3], gameboard[i+2][j+3], gameboard[i+3][j+3]}

			diagone := []rune{gameboard[i][j], gameboard[i+1][j+1], gameboard[i+2][j+2], gameboard[i+3][j+3]}
			diagtwo := []rune{gameboard[i][j+3], gameboard[i+1][j+2], gameboard[i+2][j+1], gameboard[i+3][j]}

			if !slices.Contains(rowsChecked, i) {
				if string(rowone) == "XMAS" || string(rowone) == "SAMX" {
					xmasCount += 1
				}
				if string(colone) == "XMAS" || string(colone) == "SAMX" {
					xmasCount += 1
				}
				if string(coltwo) == "XMAS" || string(coltwo) == "SAMX" {
					xmasCount += 1
				}
				if string(colthree) == "XMAS" || string(colthree) == "SAMX" {
					xmasCount += 1
				}
				if string(colfour) == "XMAS" || string(colfour) == "SAMX" {
					xmasCount += 1
				}
				rowsChecked = append(rowsChecked, i)
			}
			if !slices.Contains(rowsChecked, i+1) {
				if string(rowtwo) == "XMAS" || string(rowtwo) == "SAMX" {
					xmasCount += 1
				}
				rowsChecked = append(rowsChecked, i+1)
			}
			if !slices.Contains(rowsChecked, i+2) {
				if string(rowthree) == "XMAS" || string(rowthree) == "SAMX" {
					xmasCount += 1
				}
				rowsChecked = append(rowsChecked, i+2)
			}
			if !slices.Contains(rowsChecked, i+3) {
				if string(rowfour) == "XMAS" || string(rowfour) == "SAMX" {
					xmasCount += 1
				}
				rowsChecked = append(rowsChecked, i+3)
			}

			if string(diagone) == "XMAS" || string(diagone) == "SAMX" {
				xmasCount += 1
			}
			if string(diagtwo) == "XMAS" || string(diagtwo) == "SAMX" {
				xmasCount += 1
			}

			if slices.Contains(rowsChecked, 139) {
				fmt.Println(rowfour)
			}
		}
	}
	results[0] = xmasCount
	fmt.Println(rowsChecked)

	return results
}
