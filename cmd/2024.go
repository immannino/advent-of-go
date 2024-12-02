package cmd

import (
	"advent-of-code/internal"
	"advent-of-code/internal/data"
	"slices"
	"strconv"
	"strings"
)

func NewYear2024() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1: Historian Hysteria", Day1_2024),
	}

	return internal.Year{
		Title:   "Year 2024",
		Puzzles: Days,
	}
}

type Day12024 struct {
	input string
}

func Day1_2024() internal.Answer {
	w := Day12024{}
	w.input = data.ReadAsString("data/2024/day1.txt")

	return internal.Answer{Part1: strconv.Itoa(w.part1()), Part2: strconv.Itoa(w.part2())}
}

func (w *Day12024) part1() int {
	rows := strings.Split(w.input, "\n")

	left := make([]int, len(rows))
	right := make([]int, len(rows))

	for i, r := range rows {
		parts := strings.Split(r, "   ")
		left[i], _ = strconv.Atoi(parts[0])
		right[i], _ = strconv.Atoi(parts[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			sum += left[i] - right[i]
		} else {
			sum += right[i] - left[i]
		}
	}

	return sum
}

func (w *Day12024) part2() int {
	rows := strings.Split(w.input, "\n")

	left := make([]int, len(rows))
	right := make(map[int]int)

	for i, r := range rows {
		parts := strings.Split(r, "   ")
		left[i], _ = strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		if _, ok := right[r]; ok {
			right[r] += 1
		} else {
			right[r] = 1
		}
	}

	sum := 0

	for _, v := range left {
		if _, ok := right[v]; ok {
			sum += (v * right[v])
		}
	}

	return sum
}
