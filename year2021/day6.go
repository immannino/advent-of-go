package year2021

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

// --- Day 6: Lanternfish ---
func Day6() {
	data := utils.GetData("data/2021/day6.txt")
	fishes := make([]int, 9)

	for _, v := range strings.Split(data, ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		fishes[n] += 1
	}

	part1 := 0
	for day := 0; day < 256; day++ {
		tempFishes := make([]int, 9)

		if day == 80 {
			for k := range fishes {
				part1 += fishes[k]
			}
		}

		for k, v := range fishes {
			if k == 0 {
				tempFishes[6] += v
				tempFishes[8] += v
			} else {
				tempFishes[k-1] += v
			}
		}

		copy(fishes, tempFishes)
	}

	part2 := 0
	for _, v := range fishes {
		part2 += v
	}

	fmt.Printf("--- Day 6: Lanternfish --- { 1: %d, 2: %d }\n", part1, part2)
}
