package cmd

import (
	"advent-of-code/internal"
	"advent-of-code/internal/data"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func NewYear2023() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1: Trebuchet?!", Day1_2023),
		internal.NewPuzzle(2, "Day 2: Cube Conundrum", Day2_2023),
		internal.NewPuzzle(3, "Day 3: Gear Ratios", Day3_2023),
		internal.NewPuzzle(4, "Day 4: Scratchcards", Day4_2023),
		internal.NewPuzzle(5, "Day 5: If You Give A Seed A Fertilizer", Day5_2023),
	}

	return internal.Year{"Year 2023", Days}
}

////////////////////////////////////////////////////
/// Day 1
////////////////////////////////////////////////////

// Optional state for Day1
type Day12023 struct {
	input     []string
	intString string
	namedInts map[string]string
}

func Day1_2023() internal.Answer {
	w := Day12023{}
	input := data.ReadAsString("data/2023/day1.txt")
	w.input = strings.Split(input, "\n")
	w.intString = "1234567890"
	w.namedInts = map[string]string{
		"one":   "o1one",
		"two":   "t2two",
		"three": "t3three",
		"four":  "f4four",
		"five":  "f5five",
		"six":   "s6six",
		"seven": "s7seven",
		"eight": "e8eight",
		"nine":  "n9nine",
		"zero":  "z0zero",
	}

	var part1 int
	// Part 1
	{
		for _, v := range w.input {
			part1 += w.getFirstLastInt(v)
		}
	}

	var part2 int
	// Part 2
	{
		for _, v := range w.input {
			part2 += w.getFirstLastNamed(v)
		}
	}

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func (w *Day12023) getFirstLastInt(s string) int {
	var first, last string

	for _, v := range s {
		if strings.Contains(w.intString, string(v)) {
			if len(first) == 0 {
				first = string(v)
			} else {
				last = string(v)
			}
		}
	}

	if len(last) == 0 {
		last = first
	}

	num, _ := strconv.Atoi(first + last)
	return num
}

func (w *Day12023) getFirstLastNamed(s string) int {
	var first, last string

	val := s
	for k, r := range w.namedInts {
		val = strings.ReplaceAll(val, k, r)
	}

	for _, v := range val {
		if strings.Contains(w.intString, string(v)) {
			if len(first) == 0 {
				first = string(v)
			} else {
				last = string(v)
			}
		}
	}

	if len(last) == 0 {
		last = first
	}

	num, _ := strconv.Atoi(first + last)
	return num
}

////////////////////////////////////////////////////
/// Day 2
////////////////////////////////////////////////////

type Day22023 struct {
	input    []string
	gameMap  map[string]map[string]int
	part1Req map[string]int
}

func Day2_2023() internal.Answer {
	w := Day22023{}
	input := data.ReadAsString("data/2023/day2.txt")
	w.input = strings.Split(input, "\n")
	w.part1Req = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	w.parseInput()

	part1 := 0
	part2 := 0
	for i := 1; i <= 100; i++ {
		k := fmt.Sprintf("Game %d", i)
		v := w.gameMap[k]

		fmt.Println(k, "Red:", v["red"], "Green:", v["green"], "Blue:", v["blue"])

		// Calculate winning hand
		if len(v) == 3 && v["red"] <= w.part1Req["red"] && v["green"] <= w.part1Req["green"] && v["blue"] <= w.part1Req["blue"] {
			gameParts := strings.Split(k, "Game ")
			num, _ := strconv.Atoi(gameParts[1])
			part1 += num
		}

		// Calculate power
		part2 += v["red"] * v["green"] * v["blue"]
	}

	return internal.Answer{Part1: strconv.Itoa(part1), Part2: strconv.Itoa(part2)}
}

func (w *Day22023) parseInput() {
	w.gameMap = make(map[string]map[string]int)

	for _, v := range w.input {
		game := strings.Split(v, ": ")
		w.gameMap[game[0]] = make(map[string]int)
		pulls := strings.Split(game[1], "; ")

		for _, g := range pulls {
			cubes := strings.Split(g, ", ")

			for _, c := range cubes {
				cubeParts := strings.Split(c, " ")
				cubeCount, _ := strconv.Atoi(cubeParts[0])
				if _, ok := w.gameMap[game[0]][cubeParts[1]]; !ok {
					w.gameMap[game[0]][cubeParts[1]] = cubeCount
				} else {
					if cubeCount > w.gameMap[game[0]][cubeParts[1]] {
						w.gameMap[game[0]][cubeParts[1]] = cubeCount
					}
				}
			}
		}
	}
}

////////////////////////////////////////////////////
/// Day 3
////////////////////////////////////////////////////

type Day32023 struct {
	input []string
	part1 int
	part2 int
}

func Day3_2023() internal.Answer {
	w := Day32023{}
	input := data.ReadAsString("data/2023/day3.txt")
	w.input = strings.Split(input, "\n")

	adjacentNums := []string{}
	for y, l := range w.input {
		// a map of num chars & isAdjacentToSymbol
		currNumAdjacencies := make(map[string]bool)
		currNum := ""
		for x, c := range l {
			key := strconv.Itoa(x) + string(c)
			if c >= 48 && c <= 57 {
				currNum += string(c)
				if _, ok := currNumAdjacencies[key]; !ok {
					currNumAdjacencies[key] = false
				}

				if w.isAdjacentToSymbol(x, y, len(l)) {
					currNumAdjacencies[key] = true
				}

				continue
			}

			// We only get here at the end of a cycle
			if len(currNum) > 0 {
				if w.hasAdjacenies(currNumAdjacencies) {
					adjacentNums = append(adjacentNums, currNum)
				}

				currNum = ""
				currNumAdjacencies = make(map[string]bool)
			}
		}
	}

	for _, v := range adjacentNums {
		num, _ := strconv.Atoi(v)
		w.part1 += num
	}

	return internal.Answer{Part1: strconv.Itoa(w.part1)}
}

func (w *Day32023) hasAdjacenies(m map[string]bool) bool {
	for _, v := range m {
		if v {
			return true
		}
	}

	return false
}

func (w *Day32023) isAdjacentToSymbol(x, y, lineLen int) bool {
	// Top
	if y > 0 && w.isSymbol(rune(w.input[y-1][x])) {
		// log.Print("Symbol Top!")
		return true
	}

	// Bottom
	if y < (len(w.input)-1) && w.isSymbol(rune(w.input[y+1][x])) {
		// log.Print("Symbol Bottom!")
		return true
	}

	// Left
	if x > 0 && w.isSymbol(rune(w.input[y][x-1])) {
		// log.Print("Symbol Left!")
		return true
	}

	// Right
	if x < (lineLen-1) && w.isSymbol(rune(w.input[y][x+1])) {
		// log.Print("Symbol Right!")
		return true
	}

	// Top Right
	if (y > 0) && (x < (lineLen - 1)) && w.isSymbol(rune(w.input[y-1][x+1])) {
		// log.Print("Symbol Top Right!")
		return true
	}

	// Top Left
	if (y > 0) && (x > 0) && w.isSymbol(rune(w.input[y-1][x-1])) {
		// log.Print("Symbol Top Left!")
		return true
	}

	if y < (len(w.input)-1) && x < (lineLen-1) && w.isSymbol(rune(w.input[y+1][x+1])) {
		// log.Print("Symbol Bottom Right!")
		return true
	}

	if y < (len(w.input)-1) && x > 0 && w.isSymbol(rune(w.input[y+1][x-1])) {
		// log.Print("Symbol Bottom Right!")
		return true
	}

	return false

}

func (w *Day32023) isSymbol(c rune) bool {
	// check . or 0-9
	if c == 46 || (c >= 48 && c <= 57) {
		return false
	}

	return true
}

////////////////////////////////////////////////////
/// Day 4
////////////////////////////////////////////////////

type Day42023 struct {
	input []string
	part1 int
	part2 int
}

func Day4_2023() internal.Answer {
	w := Day42023{}
	input := data.ReadAsString("data/2023/day4.txt")
	w.input = strings.Split(input, "\n")
	w.part1 = 0
	w.part2 = 0

	gameCount := make(map[string]int)

	for i, line := range w.input {
		gameParts := strings.Split(line, ": ")
		if _, ok := gameCount[strconv.Itoa(i)]; !ok {
			// first game starts as 1
			gameCount[strconv.Itoa(i)] = 1
		} else {
			// additional games append
			gameCount[strconv.Itoa(i)] += 1
		}

		numberParts := strings.Split(gameParts[1], " | ")
		winning := strings.Split(numberParts[0], " ")
		elfNumbers := strings.Split(numberParts[1], " ")
		winCount := 0
		for _, g := range winning {
			if g == "" || g == " " {
				continue
			}
			for _, e := range elfNumbers {
				if e == "" || e == " " {
					continue
				}
				if g == e {
					if winCount == 0 {
						winCount = 1
					} else {
						winCount += 1
					}
				}
			}
		}

		if winCount > 0 {
			w.part1 += int(math.Exp2(float64(winCount) - 1))
		}

		for g := 1; g <= winCount; g++ {
			if _, ok := gameCount[strconv.Itoa(g+i)]; !ok {
				gameCount[strconv.Itoa(g+i)] = (1 * gameCount[strconv.Itoa(i)])
			} else {
				gameCount[strconv.Itoa(g+i)] += (1 * gameCount[strconv.Itoa(i)])
			}
		}
	}

	for _, g := range gameCount {
		w.part2 += g
	}

	return internal.Answer{Part1: strconv.Itoa(w.part1), Part2: strconv.Itoa(w.part2)}
}

////////////////////////////////////////////////////
/// Day 5
////////////////////////////////////////////////////

type Day52023 struct {
	input []string
	part1 int
	part2 int
}

func Day5_2023() internal.Answer {
	return internal.Answer{}
}
