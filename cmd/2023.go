package cmd

import (
	"advent-of-code/internal"
	"advent-of-code/internal/data"
	"strconv"
	"strings"
)

func NewYear2023() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1: Trebuchet?!", Day1_2023),
	}

	return internal.Year{"Year 2023", Days}
}

// Optional state for Day1
type Day12023 struct {
	input     []string
	intString string
	namedInts map[string]string
}

func Day1_2023() (string, string) {
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

	return strconv.Itoa(part1), strconv.Itoa(part2)
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
