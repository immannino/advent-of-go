package year2022

import (
	"advent-of-code/pkg/data"
	"fmt"
	"strings"
)

func Day3() string {
	var part1 int
	var part2 int

	// Index of letter is priority value numerical - e.g. alphabet['k'] == index + 1 priority
	var alphabet string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`

	// 	var example string = `vJrwpWtwJgWrhcsFMMfFFhFp
	// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	// PmmdzqPrVvPwwTWBwg
	// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	// ttgJtRGJQctTZtZT
	// CrZsJsPPZsGzwwsLwLmpwMDw
	// `
	input := data.ReadAsString("data/2022/day3.txt")
	compartments := strings.Split(input, "\n")

	// Part 1
	for _, c := range compartments {
		left := c[:len(c)/2] // left half of the compartment to compare against
		right := c[len(c)/2:]

		for _, p := range right {
			if strings.Contains(left, string(p)) {
				part1 += strings.Index(alphabet, string(p)) + 1
				break
			}
		}
	}

	// Part 2 - Segments of 3
	for i := 0; i < len(compartments)/3; i++ {
		segs := compartments[i*3 : (i*3)+3] // Groups of 3 inputs

		for _, p := range segs[0] {
			if strings.Contains(segs[1], string(p)) && strings.Contains(segs[2], string(p)) {
				part2 += strings.Index(alphabet, string(p)) + 1
				break
			}
		}
	}

	return fmt.Sprintf("---| Day 3 Rocksack Reorganization - 1: %d 2: %d |---\n", part1, part2)

}
