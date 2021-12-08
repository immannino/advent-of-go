package year2021

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

// --- Day 7: The Treachery of Whales ---
func Day7() {
	data := utils.GetData("data/2021/day7.txt")
	rows := strings.Split(data, ",")
	crabs, min, max := getCrabs(rows)
	horizontals := make([]int, (min+max)+1)
	horizontalsMultiplied := make([]int, (min+max)+1)

	for i := min; i <= max; i++ {
		diff := 0
		diff2 := 0

		for _, v := range crabs {
			if v > i {
				diff += (v - i)
				diff2 += summation((v - i))
			} else {
				diff += (i - v)
				diff2 += summation((i - v))
			}
		}

		horizontals[i] = diff
		horizontalsMultiplied[i] = diff2
	}

	smallest := 0
	smallestMult := 0

	for i := 0; i < len(horizontals); i++ {
		if horizontals[i] < horizontals[smallest] {
			smallest = i
		}

		if horizontalsMultiplied[i] < horizontalsMultiplied[smallestMult] {
			smallestMult = i
		}
	}

	part1 := horizontals[smallest]
	part2 := horizontalsMultiplied[smallestMult]

	fmt.Printf("--- Day 7: The Treachery of Whales --- { 1: %d, 2: %d }\n", part1, part2)
}

func summation(n int) int {
	if n%2 == 0 {
		return (n / 2) * (n + 1)
	}

	return ((n + 1) / 2) * n
}

func getCrabs(rows []string) ([]int, int, int) {
	nums := []int{}

	n, err := strconv.Atoi(rows[0])
	if err != nil {
		panic(err)
	}

	nums = append(nums, n)
	min := n
	max := n

	for i := 1; i < len(rows); i++ {
		n, err := strconv.Atoi(rows[i])
		if err != nil {
			panic(err)
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}

		nums = append(nums, n)
	}

	return nums, min, max
}
