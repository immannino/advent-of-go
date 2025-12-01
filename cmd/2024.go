package cmd

import (
	"advent-of-code/internal"
	"advent-of-code/internal/data"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func NewYear2024() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1: Historian Hysteria", Day1_2024),
		internal.NewPuzzle(1, "Day 2: Red-Nosed Reports", Day2_2024),
	}

	return internal.Year{
		Title:   "Year 2024",
		Puzzles: Days,
	}
}

type Day12024 struct {
	input string
}

func Day1_2024() internal.Answer {
	w := Day12024{}
	w.input = data.ReadAsString("data/2024/day1.txt")

	return internal.Answer{Part1: strconv.Itoa(w.part1()), Part2: strconv.Itoa(w.part2())}
}

func (w *Day12024) part1() int {
	rows := strings.Split(w.input, "\n")

	left := make([]int, len(rows))
	right := make([]int, len(rows))

	for i, r := range rows {
		parts := strings.Split(r, "   ")
		left[i], _ = strconv.Atoi(parts[0])
		right[i], _ = strconv.Atoi(parts[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			sum += left[i] - right[i]
		} else {
			sum += right[i] - left[i]
		}
	}

	return sum
}

func (w *Day12024) part2() int {
	rows := strings.Split(w.input, "\n")

	left := make([]int, len(rows))
	right := make(map[int]int)

	for i, r := range rows {
		parts := strings.Split(r, "   ")
		left[i], _ = strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		if _, ok := right[r]; ok {
			right[r] += 1
		} else {
			right[r] = 1
		}
	}

	sum := 0

	for _, v := range left {
		if _, ok := right[v]; ok {
			sum += (v * right[v])
		}
	}

	return sum
}

type Day22024 struct {
	input string
}

func Day2_2024() internal.Answer {
	w := Day22024{}
	w.input = data.ReadAsString("data/2024/day2_test.txt")

	return internal.Answer{Part1: strconv.Itoa(w.part1()), Part2: strconv.Itoa(w.part2())}
}

func (w *Day22024) part1() int {
	count := 0

	for x, r := range strings.Split(w.input, "\n") {
		elems := strings.Split(r, " ")
		var increasing, decreaing, fail bool

		for i, v := range elems {
			// fmt.Println(i, v, increasing, decreaing, fail)
			if i == 0 {
				continue
			}

			c := toInt(v)
			p := toInt(elems[i-1])

			if !w.isWithinRange(c, p) {
				fail = true
				break
			}

			if i == 1 {
				if c > p {
					increasing = true
				} else {
					decreaing = true
				}

				continue
			}

			if increasing {
				if c < p {
					fail = true
					break
				}
			}

			if decreaing {
				if c > p {
					fail = true
					break
				}
			}
		}

		if !fail {
			count += 1
			fmt.Printf("PASS: %d, %s\n", x, r)
		} else {

			fmt.Printf("FAIL: %d, %s\n", x, r)
		}

	}

	return count
}

func (w *Day22024) part2() int {

	count := 0

	for x, r := range strings.Split(w.input, "\n") {
		elems := strings.Split(r, " ")
		var increasing, decreaing, fail, fault bool

		// Change the flow to essentially duplicate computation:
		// 1. if there was a problem with a number
		// 2. Remove number from slice
		// 3. re-run a refresh flow without the number
		// If that also failed, fail, otherwise succeed
		for i, v := range elems {
			// fmt.Println(i, v, increasing, decreaing, fail)
			if i == 0 {
				continue
			}

			c := toInt(v)
			var p int
			if fault {
				p = toInt(elems[i-2])
			} else {
				p = toInt(elems[i-1])
			}

			// Always fails
			if !w.isWithinRange(c, p) {
				fail = true
				break
			}

			if i == 1 {
				if c > p {
					increasing = true
				} else {
					decreaing = true
				}

				continue
			}

			if increasing {
				if c < p {
					if !fault {
						fault = true
						continue
					} else {
						fail = true
						break
					}
				}
			}

			if decreaing {
				if c > p {
					if !fault {
						fault = true
						continue
					} else {
						fail = true
						break
					}
				}
			}
		}

		if !fail {
			count += 1
			fmt.Printf("PASS: %d, %s\n", x, r)
		} else {

			fmt.Printf("FAIL: %d, %s\n", x, r)
		}

	}

	return count
}

func toInt(e string) int {
	v, _ := strconv.Atoi(e)
	return v
}

func (w *Day22024) isWithinRange(a, b int) bool {
	var s int

	if a > b {
		s = a - b
	} else {
		s = b - a
	}

	if s >= 1 && s <= 3 {
		return true
	}
	return false
}
