package year2022

import (
	"advent-of-code/pkg/data"
	"strconv"
)

func Day6() []string {
	input := data.ReadAsString("data/2022/day6.txt")
	part1 := find(input, 4)
	part2 := find(input, 14)

	return []string{"Day 6: Tuning Trouble", strconv.Itoa(part1), strconv.Itoa(part2)}
}

func find(input string, coef int) int {
	for i := range input {
		next := input[i : i+coef]
		has := make(map[rune]bool)
		fail := false

		for _, v := range next {
			if _, ok := has[v]; ok {
				fail = true
				break
			}

			has[v] = true
		}

		if !fail {
			return i + coef
		}
	}

	return 0
}
