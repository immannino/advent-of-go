package year2021

import (
	"advent-of-code/pkg/data"
	"fmt"
	"strconv"
	"strings"
)

// --- Day 2: Dive! ---
func Day2() {
	data := data.ReadAsString("data/2021/day2.txt")
	rows := strings.Split(data, "\n")
	part1 := len(data)
	part2 := 0

	depth := 0
	depth2 := 0
	horizontal := 0
	aim := 0

	for _, v := range rows {
		if strings.Contains(v, "forward") {
			val := getNum(v, "forward ")
			horizontal += val
			depth2 += val * aim
		} else if strings.Contains(v, "up") {
			val := getNum(v, "up ")
			depth -= val
			aim -= val
		} else if strings.Contains(v, "down") {
			val := getNum(v, "down ")
			depth += val
			aim += val
		}
	}

	part1 = horizontal * depth
	part2 = horizontal * depth2

	fmt.Printf("--- Day 2: Dive! --- { 1: %d, 2: %d }\n", part1, part2)
}

func getNum(input string, sep string) int {
	numString := strings.Split(input, sep)[1]
	num, err := strconv.Atoi(numString)

	if err != nil {
		panic(err)
	}

	return num
}
