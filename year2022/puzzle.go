package year2022

import (
	table "advent-of-code/pkg/table"
	"advent-of-code/pkg/util"
)

func Solve() {
	v := [][]string{
		util.TimeAppend(Day1),
		util.TimeAppend(Day2),
		util.TimeAppend(Day3),
		util.TimeAppend(Day4),
		util.TimeAppend(Day5),
		util.TimeAppend(Day6),
		util.TimeAppend(Day7),
	}

	table.Print(v)
}
