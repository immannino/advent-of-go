package year2022

import (
	"advent-of-code/pkg/data"
	"fmt"
	"strconv"
	"strings"
)

func Day4() string {
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
		a := expand(parse(parts[0]))
		b := expand(parse(parts[1]))

		if strings.Contains(a, b) || strings.Contains(b, a) {
			part1 += 1
		}
	}
	return fmt.Sprintf("---| Day 4 Camp Cleanup - 1: %d 2: %d |---\n", part1, part2)

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

func expand(b, e int) string {
	var s string

	for i := b; i <= e; i++ {
		s += fmt.Sprintf("%s,", strconv.Itoa(i))
	}

	return s
}
