package cmd

import (
	"advent-of-code/internal"
	"advent-of-code/internal/data"
	"advent-of-code/internal/util"
	"strconv"
	"strings"
)

func NewYear2025() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1: Secret Entrance", Day1_2025),
	}

	return internal.Year{
		Title:   "Year 2025",
		Puzzles: Days,
	}
}

type Day12025 struct {
	input string
}

func Day1_2025() internal.Answer {
	w := Day12025{}
	w.input = data.ReadAsString("data/2025/day1.txt")

	return internal.Answer{Part1: strconv.Itoa(w.part1()), Part2: strconv.Itoa(w.part2())}
}

func (w *Day12025) part1() int {
	cmds := strings.Split(w.input, "\n")

	dial := 50
	z := 0
	for _, c := range cmds {
		action := c[0]
		d := util.MustInt(c[1:])
		switch string(action) {
		case "R":
			dial = (dial + d) % 100
		case "L":
			dial = (dial - d) % 100
		}

		if dial == 0 {
			z += 1
		}
	}

	return z
}

func (w *Day12025) part2() int {
	cmds := strings.Split(w.input, "\n")

	dial := 50
	r := 0
	for _, c := range cmds {
		action := c[0]
		d := util.MustInt(c[1:])

		mod := 1
		if string(action) == "L" {
			mod = -1
		}

		// Part 1
		for clicks := 0; clicks < d; clicks++ {
			dial += mod

			if dial == 0 {
				r += 1
			}

			if dial < 0 {
				dial = 99
			}

			if dial > 99 {
				dial = 0
				r += 1
			}
		}

		// roundRations := d / 99
		// r += roundRations
		// switch string(action) {
		// case "R":
		// 	if roundRations == 0 {
		// 		if (dial + d) > 99 {
		// 			r += 1
		// 		}
		// 	}
		// 	dial = (dial + d) % 100
		// case "L":
		// 	if roundRations == 0 {
		// 		if (dial - d) < 0 {
		// 			r += 1
		// 		}
		// 	}
		// 	dial = mod((dial - d), 100)
		// }

		// if dial == 0 {
		// 	z += 1
		// 	r += 1
		// }

		// fmt.Println(dial, r)
		// fmt.Println()
		// if i == 10 {
		// 	break
		// }
	}

	return r
}

func mod(n int, m int) int {
	return ((n % m) + m) % m
}
