package advent

import (
	"fmt"
	"github.com/one-cold-canuck/advent2024/internal/util"
	"strconv"
	"strings"
)

var compass = map[rune][]int{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

var routemap [][]byte
var startpos []int = []int{0, 0}
var direction rune
var currpos []int = []int{0, 0}
var visitedMap map[string]string

func RunDaySix() []int {
	results := []int{0, 0}

	InitializeRoutemap("data/day6-map.txt")

	if len(startpos) < 1 {
		fmt.Println("No start position found!")
		return results
	}

	direction = rune(routemap[startpos[0]][startpos[1]])

	onmap := true
	currpos = startpos
	visitedMap = make(map[string]string)
	loopcount := 0

	for onmap {
		currposkey := strconv.Itoa(currpos[0]) + "," + strconv.Itoa(currpos[1])
		_, ok := visitedMap[currposkey]
		if ok {
			if visitedMap[currposkey] == string(GetNextDirection(direction)) {
				loopcount += 1
				for k := range visitedMap {
					delete(visitedMap, k)
				}
			}
		}
		onmap = MoveGuard(currpos, direction)
	}

	for _, v := range routemap {
		results[0] += strings.Count(string(v), "X") + strings.Count(string(v), "^") + strings.Count(string(v), ">") + strings.Count(string(v), "v") + strings.Count(string(v), "<")
		fmt.Printf("%s\n", v)
	}
	results[1] = loopcount
	return results
}

func GetNextDirection(dir rune) rune {
	var nextdir rune
	switch dir {
	case '^':
		nextdir = '>'
		break
	case '>':
		nextdir = 'v'
		break
	case 'v':
		nextdir = '<'
		break
	case '<':
		nextdir = '^'
		break
	default:
	}
	return nextdir
}

func MoveGuard(cpos []int, dir rune) bool {

	nextY := cpos[0] + compass[dir][0]
	nextX := cpos[1] + compass[dir][1]
	coordkey := strconv.Itoa(cpos[0]) + "," + strconv.Itoa(cpos[1])
	visitedMap[coordkey] = string(dir)

	if !((nextY >= len(routemap) || nextY < 0) || (nextX >= len(routemap[cpos[0]]) || nextX < 0)) {
		if routemap[nextY][nextX] == '.' || routemap[nextY][nextX] == 'X' || routemap[nextY][nextX] == '^' || routemap[nextY][nextX] == '>' || routemap[nextY][nextX] == 'v' || routemap[nextY][nextX] == '<' {
			routemap[cpos[0]][cpos[1]] = byte(dir)
			// routemap[nextY][nextX] = byte(dir)
			cpos[0] = nextY
			cpos[1] = nextX
			currpos = cpos
			return true
		} else if routemap[nextY][nextX] == '#' {
			switch dir {
			case '^':
				dir = '>'
				break
			case '>':
				dir = 'v'
				break
			case 'v':
				dir = '<'
				break
			case '<':
				dir = '^'
				break
			default:
			}

			// routemap[cpos[0]][cpos[1]] = byte(dir)
			direction = dir
			currpos = cpos

			return true
		} else {
			fmt.Println("Undefined movement!")
		}
	}
	routemap[cpos[0]][cpos[1]] = 'X'
	return false
}

func InitializeRoutemap(filename string) {

	floorplan := util.ReadFileIntoStrings(filename, "\n")
	floorplan = floorplan[:len(floorplan)-1]
	for i, tileRow := range floorplan {
		routemap = append(routemap, []byte(tileRow))
		if strings.Index(tileRow, "^") > -1 {
			startpos[0] = i
			startpos[1] = strings.Index(tileRow, "^")
		} else if strings.Index(tileRow, "<") > -1 {
			startpos[0] = i
			startpos[1] = strings.Index(tileRow, "<")
		} else if strings.Index(tileRow, ">") > -1 {
			startpos[0] = i
			startpos[1] = strings.Index(tileRow, "<")
		} else if strings.Index(tileRow, "v") > -1 {
			startpos[0] = i
			startpos[1] = strings.Index(tileRow, "<")
		}
	}
}
