package year2022

import (
	"advent-of-code/pkg/data"
	"advent-of-code/pkg/util"
	"fmt"
	"strconv"
	"strings"
)

func Day5() string {
	var part1 string
	var part2 string

	input := data.ReadAsString("data/2022/day5.txt")
	instructions := strings.Split(input, "\n\n")
	stacks, cols := parseBoxes(instructions[0])
	stacksPart2 := make([]string, cols)

	copy(stacksPart2, stacks)

	actions := strings.Split(instructions[1], "\n")

	fmt.Println(len(stacks), len(stacksPart2), len(actions))
	for _, line := range actions {
		// Part 1
		count, from, to := parseSupplyLine(line)
		// Find items to move
		sub := stacks[from][len(stacks[from])-count:]
		// Find new state of From
		newFrom := stacks[from][:len(stacks[from])-count]

		// Set new From
		stacks[from] = newFrom
		// Append to To
		stacks[to] += util.Reverse(sub)

		// Part 2 (the same, minus util.Reverse)
		sub2 := stacksPart2[from][len(stacksPart2[from])-count:]
		newFrom2 := stacksPart2[from][:len(stacksPart2[from])-count]
		stacksPart2[from] = newFrom2
		stacksPart2[to] += string(sub2)
	}

	for _, v := range stacks {
		if len(v) > 0 {
			part1 += string(v[len(v)-1])
		}
	}

	for _, v := range stacksPart2 {
		if len(v) > 0 {
			part2 += string(v[len(v)-1])
		}
	}

	return fmt.Sprintf("---| Day 5: Supply Stacks - 1: %s 2: %v |---\n", part1, part2)

}

func print(s []string) {
	for i, v := range s {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func parseBoxes(input string) ([]string, int) {
	lines := strings.Split(input, "\n")
	linesWithout := lines[:len(lines)-1]
	cols := (len(lines[0]) / 4) + 1
	boxes := make([]string, cols)

	for _, l := range linesWithout {
		l += " " // Pad a space
		for i := 0; i < cols; i++ {
			item := l[i*4 : (i*4)+4]
			if string(item[0]) == "[" {
				val := strings.Trim(item, " []")
				boxes[i] += val
			}
		}
	}

	for i := range boxes {
		boxes[i] = util.Reverse(boxes[i])
	}

	return boxes, cols
}

func parseSupplyLine(i string) (int, int, int) {
	parts := strings.Split(strings.Replace(
		strings.Replace(
			strings.Replace(i, "move ", "", 1),
			" from ",
			",",
			1),
		" to ",
		",",
		1), ",")

	count, err := strconv.ParseInt(parts[0], 10, 16)
	if err != nil {
		panic(err)
	}

	from, err := strconv.ParseInt(parts[1], 10, 16)
	if err != nil {
		panic(err)
	}

	to, err := strconv.ParseInt(parts[2], 10, 16)
	if err != nil {
		panic(err)
	}

	return int(count), int(from) - 1, int(to) - 1

}
