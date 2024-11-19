package cmd

import (
	"advent-of-code/internal"
	"fmt"
)

func NewYear2024() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1: <TBD>", Day1_2024),
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

	fmt.Println(w)

	return internal.Answer{}
}
