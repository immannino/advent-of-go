package year2015

import (
	"advent-of-code/utils"
	"fmt"
)

// Day1 -- Not Quite Lisp
func Day1() {
	data := utils.GetData("data/2015/day1.txt")

	floor := 0
	input := []rune(data)
	position := 1
	hasMatch := false

	for _, e := range input {
		if e == '(' {
			floor += 1
		} else {
			floor -= 1
		}

		if !hasMatch && floor == -1 {
			hasMatch = true
		}

		if !hasMatch {
			position += 1
		}
	}

	fmt.Printf("Day 1: { 1: %d, 2: %d }\n", floor, position)
}
