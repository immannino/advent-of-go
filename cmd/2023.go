package cmd

import (
	"advent-of-code/internal"
)

func NewYear2023() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1", Day1_2023),
	}

	return internal.Year{"Year 2023", Days}
}

func Day1_2023() (string, string) {
	return "", ""
}
