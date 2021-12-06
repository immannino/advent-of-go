package year2021

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

// --- Day 5: Hydrothermal Venture ---
func Day5() {
	data := utils.GetData("data/2021/day5.txt")
	rows := strings.Split(data, "\n")
	collisions := make(map[string]int)
	collisions2 := make(map[string]int)
	part1 := 0

	for _, line := range rows {
		x1, y1, x2, y2 := getCoords(line)
		lStart := 0
		lEnd := 0

		if x1 == x2 {
			lStart = y1
			lEnd = y2

			if y1 > y2 {
				lStart = y2
				lEnd = y1
			}

			for i := lStart; i <= lEnd; i++ {
				key := fmt.Sprintf("%d,%d", x1, i)

				if _, ok := collisions[key]; ok {
					collisions[key] += 1
				} else {
					collisions[key] = 1
				}

				if _, ok := collisions2[key]; ok {
					collisions2[key] += 1
				} else {
					collisions2[key] = 1
				}
			}
		} else if y1 == y2 {
			lStart = x1
			lEnd = x2

			if x1 > x2 {
				lStart = x2
				lEnd = x1
			}

			for i := lStart; i <= lEnd; i++ {
				key := fmt.Sprintf("%d,%d", i, y1)

				if _, ok := collisions[key]; ok {
					collisions[key] += 1
				} else {
					collisions[key] = 1
				}

				if _, ok := collisions2[key]; ok {
					collisions2[key] += 1
				} else {
					collisions2[key] = 1
				}
			}
		} else { // Diagonally
			y := y1

			if x1 < x2 { // R -> L
				for x := x1; x <= x2; x++ {
					key := fmt.Sprintf("%d,%d", x, y)

					if _, ok := collisions2[key]; ok {
						collisions2[key] += 1
					} else {
						collisions2[key] = 1
					}

					if y < y2 {
						y += 1
					} else {
						y -= 1
					}
				}

			} else { // L -> R (backwards)
				for x := x1; x >= x2; x-- {
					key := fmt.Sprintf("%d,%d", x, y)

					if _, ok := collisions2[key]; ok {
						collisions2[key] += 1
					} else {
						collisions2[key] = 1
					}

					if y < y2 {
						y += 1
					} else {
						y -= 1
					}
				}
			}
		}
	}

	part2 := 0

	for _, v := range collisions {
		if v > 1 {
			part1 += 1
		}
	}

	for _, v := range collisions2 {
		if v > 1 {
			part2 += 1
		}
	}

	fmt.Printf("--- Day 5: Hydrothermal Venture --- { 1: %d, 2: %d }\n", part1, part2)
}

func getCoords(s string) (int, int, int, int) {
	parts := strings.Split(s, " -> ")
	start := strings.Split(parts[0], ",")
	end := strings.Split(parts[1], ",")

	x1, err := strconv.Atoi(start[0])
	if err != nil {
		panic(err)
	}

	y1, err := strconv.Atoi(start[1])
	if err != nil {
		panic(err)
	}

	x2, err := strconv.Atoi(end[0])
	if err != nil {
		panic(err)
	}

	y2, err := strconv.Atoi(end[1])
	if err != nil {
		panic(err)
	}

	return x1, y1, x2, y2
}
