package year2021

import (
	"advent-of-code/pkg/data"
	"fmt"
	"strconv"
	"strings"
)

// --- Day 1: Sonar Sweep ---
func Day1() {
	data := data.ReadAsString("data/2021/day1.txt")
	rows := strings.Split(data, "\n")
	part1 := -1
	previous := 0

	// part 1
	for _, v := range rows {
		i, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		if i > previous {
			part1 += 1
		}

		previous = i
	}

	part2 := -1
	previous = 0

	for i := 0; i < len(rows)-2; i++ {
		m := getTMSWSum(i, rows)

		if m > previous {
			part2 += 1
		}

		previous = m
	}

	fmt.Printf("--- Day 1: Sonar Sweep --- { 1: %d, 2: %d }\n", part1, part2)
}

func getTMSWSum(start int, rows []string) int {
	sum := 0

	for i := start; i < start+3; i++ {
		num, err := strconv.Atoi(rows[i])

		if err != nil {
			panic(err)
		}

		sum += num
	}

	return sum
}
