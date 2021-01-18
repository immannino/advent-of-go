package year2015

import (
	"advent-of-code/utils"
	"fmt"
)

var data string

func init() {
	data = utils.GetData("year2015/day1.txt")
}

func Day1() {
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
