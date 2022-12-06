package year2022

import (
	"advent-of-code/pkg/data"
	"fmt"
	"strconv"
	"strings"
)

func Day4() []string {
	var part1 int
	var part2 int
	// 	var example string = `2-4,6-8
	// 2-3,4-5
	// 5-7,7-9
	// 2-8,3-7
	// 6-6,4-6
	// 2-6,4-8`
	input := data.ReadAsString("data/2022/day4.txt")
	coords := strings.Split(input, "\n")

	for _, c := range coords {
		// time.Sleep(time.Second * 3)
		parts := strings.Split(c, ",")

		aStart, aEnd := parse(parts[0])
		bStart, bEnd := parse(parts[1])

		// Part 1 comparisons
		if bStart >= aStart && bEnd <= aEnd {
			part1 += 1
			continue
		}

		if aStart >= bStart && aEnd <= bEnd {
			part1 += 1
		}
	}

	for _, c := range coords {
		parts := strings.Split(c, ",")
		aStart, aEnd := parse(parts[0])
		bStart, bEnd := parse(parts[1])

		m := make(map[int]bool)

		for i := aStart; i <= aEnd; i++ {
			m[i] = true
		}

		for i := bStart; i <= bEnd; i++ {
			if _, exists := m[i]; exists {
				part2 += 1
				break
			}
		}
	}

	return []string{"Day 4: Camp Cleanup", strconv.Itoa(part1), strconv.Itoa(part2)}

}

func parse(s string) (int, int) {
	parts := strings.Split(s, "-")
	a, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		panic(err)
	}

	b, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		panic(err)
	}

	if s != fmt.Sprintf("%d-%d", a, b) {
		panic("parse error")
	}

	return int(a), int(b)
}
