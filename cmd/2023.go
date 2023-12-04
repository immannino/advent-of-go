package cmd

import (
	"advent-of-code/internal"
	"advent-of-code/internal/data"
	"fmt"
	"strconv"
	"strings"
)

func NewYear2023() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1: Trebuchet?!", Day1_2023),
		internal.NewPuzzle(2, "Day 2: Cube Conundrum", Day2_2023),
	}

	return internal.Year{"Year 2023", Days}
}

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
