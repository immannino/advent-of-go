package year2015

import (
	"advent-of-code/pkg/data"
	"fmt"
)

// Day3 -- Perfectly Spherical Houses in a Vacuum
func Day3() {
	data := data.ReadAsString("data/2015/day3.txt")
	coords := []rune(data)
	visited := make(map[string]bool)
	pos := []int{0, 0}

	visited["{0,0}"] = true

	// Part 1
	for _, c := range coords {
		switch c {
		case '^':
			pos[1] += 1
		case '>':
			pos[0] += 1
		case 'v':
			pos[1] -= 1
		case '<':
			pos[0] -= 1
		}

		coordString := fmt.Sprintf("{%d,%d}", pos[0], pos[1])

		if !visited[coordString] {
			visited[coordString] = true
		}
	}

	// Part 2
	roboVisited := make(map[string]bool)
	robo := []int{0, 0}
	santa := []int{0, 0}

	roboVisited["{0,0}"] = true

	for i := 0; i < len(coords); i += 2 {
		santaAction := coords[i]
		roboAction := coords[i+1]

		switch santaAction {
		case '^':
			santa[1] += 1
		case '>':
			santa[0] += 1
		case 'v':
			santa[1] -= 1
		case '<':
			santa[0] -= 1
		}

		coordString := fmt.Sprintf("{%d,%d}", santa[0], santa[1])

		if !roboVisited[coordString] {
			roboVisited[coordString] = true
		}

		switch roboAction {
		case '^':
			robo[1] += 1
		case '>':
			robo[0] += 1
		case 'v':
			robo[1] -= 1
		case '<':
			robo[0] -= 1
		}

		coordString = fmt.Sprintf("{%d,%d}", robo[0], robo[1])

		if !roboVisited[coordString] {
			roboVisited[coordString] = true
		}
	}

	fmt.Printf("Day 3: { 1: %d, 2: %d }\n", len(visited), len(roboVisited))
}
