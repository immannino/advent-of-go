package year2021

import (
	"advent-of-code/pkg/data"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var sumMap = map[int]int{
	42: 0,
	17: 1,
	34: 2,
	39: 3,
	30: 4,
	37: 5,
	41: 6,
	25: 7,
	49: 8,
	45: 9,
}

// --- Day 8: Seven Segment Search ---
func Day8() {
	data := data.ReadAsString("data/2021/day8.txt")
	inputs := strings.Split(data, "\n")
	part1 := 0
	part2 := 0
	part2Values := []string{}

	for _, v := range inputs {
		parts := strings.Split(v, " | ")
		segs := strings.Split(parts[0], " ")
		keys := strings.Split(parts[1], " ")

		sort.Slice(segs, func(i, j int) bool {
			return len(segs[i]) < len(segs[j])
		})

		charMap := make(map[string]int)

		// Score the weights
		for _, v := range segs {
			digits := strings.Split(v, "")

			for _, digit := range digits {
				if _, ok := charMap[digit]; ok {
					charMap[digit] += 1
				} else {
					charMap[digit] = 1
				}
			}
		}

		tempNum := ""

		for _, k := range keys {
			if len(k) == 2 || len(k) == 4 || len(k) == 3 || len(k) == 7 {
				part1 += 1
			}

			digits := strings.Split(k, "")
			keySum := 0

			for _, d := range digits {
				keySum += charMap[d]
			}

			tempNum = fmt.Sprintf("%s%d", tempNum, sumMap[keySum])
		}

		n, err := strconv.Atoi(tempNum)
		if err != nil {
			panic(err)
		}

		part2 += n
	}

	for _, v := range part2Values {
		fmt.Println(v)
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		part2 += n
	}

	fmt.Printf("--- Day 8: Seven Segment Search --- { 1: %d, 2: %d }\n", part1, part2)
}

func sorts(s string) string {
	l := strings.Split(s, "")
	sort.Strings(l)
	return strings.Join(l, "")
}
