package year2022

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var data string
var part1 int
var part2 int
var example string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func Day1() string {
	data = utils.GetData("data/2022/day1.txt")
	// data = example
	rows := strings.Split(data, "\n")

	elves := []int{}
	currentElf := 0

	for _, r := range rows {
		if r == "" {
			elves = append(elves, currentElf)
			currentElf = 0
			continue
		}

		v, _ := strconv.ParseInt(r, 10, 32)

		currentElf += int(v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	part1 = elves[0]

	part2 = elves[0] + elves[1] + elves[2]

	return fmt.Sprintf("---| Day 1 Calorie Counting - 1: %d 2: %d |---\n", part1, part2)
}
