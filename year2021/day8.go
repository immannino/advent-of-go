package year2021

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var valueToSignal = map[int]string{
	0: "abcefg",
	1: "cf",
	2: "acdeg",
	3: "acdeg",
	4: "bcdf",
	5: "abdfg",
	6: "abdefg",
	7: "acf",
	8: "abcdefg",
	9: "abcdfg",
}

var signalToValue map[string]int

func init() {
	signalToValue = make(map[string]int)

	for k, v := range valueToSignal {
		signalToValue[v] = k
	}
}

type Digit struct {
	Top          string
	Top_left     string
	Top_right    string
	Middle       string
	Bottom       string
	Bottom_left  string
	Bottom_right string
}

// func numToString(n int, d Digit) string {
// 	switch n {
// 	case 0:
// 		return strings.Join([]string{d.Top, d.Top_left, d.Top_right, d.Bottom_left, d.Bottom_right, d.Bottom}, "")
// 	case 1:
// 		return strings.Join([]string{d.Top_right, d.Bottom_right}, "")
// 	case 2:
// 		return strings.Join([]string{d.Top, d.Top_right, d.Middle, d.Bottom_left, d.Bottom}, "")
// 	case 3:
// 		return strings.Join([]string{d.Top, d.Top_right, d.Middle, d.Bottom_right, d.Bottom}, "")
// 	case 4:
// 		return strings.Join([]string{d.Top_left, d.Top_right, d.Middle, d.Bottom_right}, "")
// 	case 5:
// 		return strings.Join([]string{d.Top, d.Top_left, d.Middle, d.Bottom_right, d.Bottom}, "")
// 	case 6:
// 		return strings.Join([]string{d.Top, d.Top_left, d.Middle, d.Bottom_left, d.Bottom_right, d.Bottom}, "")
// 	case 7:
// 		return strings.Join([]string{d.Top, d.Top_right, d.Bottom_right}, "")
// 	case 8:
// 		return strings.Join([]string{d.Top, d.Top_left, d.Top_right, d.Middle, d.Bottom_left, d.Bottom_right, d.Bottom}, "")
// 	case 9:
// 		return strings.Join([]string{d.Top, d.Top_left, d.Top_right, d.Middle, d.Bottom_right, d.Bottom}, "")
// 	}

// 	return ""
// }

// func numToString(n int, d Digit) (int, Digit) {
// 	switch n {
// 	case 0:
// 		return strings.Join([]string{d.Top, d.Top_left, d.Top_right, d.Bottom_left, d.Bottom_right, d.Bottom}, "")
// 	case 1:
// 		return strings.Join([]string{d.Top_right, d.Bottom_right}, "")
// 	case 2:
// 		return strings.Join([]string{d.Top, d.Top_right, d.Middle, d.Bottom_left, d.Bottom}, "")
// 	case 3:
// 		return strings.Join([]string{d.Top, d.Top_right, d.Middle, d.Bottom_right, d.Bottom}, "")
// 	case 4:
// 		return strings.Join([]string{d.Top_left, d.Top_right, d.Middle, d.Bottom_right}, "")
// 	case 5:
// 		return strings.Join([]string{d.Top, d.Top_left, d.Middle, d.Bottom_right, d.Bottom}, "")
// 	case 6:
// 		return strings.Join([]string{d.Top, d.Top_left, d.Middle, d.Bottom_left, d.Bottom_right, d.Bottom}, "")
// 	case 7:
// 		return strings.Join([]string{d.Top, d.Top_right, d.Bottom_right}, "")
// 	case 8:
// 		return strings.Join([]string{d.Top, d.Top_left, d.Top_right, d.Middle, d.Bottom_left, d.Bottom_right, d.Bottom}, "")
// 	case 9:
// 		return strings.Join([]string{d.Top, d.Top_left, d.Top_right, d.Middle, d.Bottom_right, d.Bottom}, "")
// 	}

// 	return ""
// }

// --- Day 8: Seven Segment Search ---
func Day8() {
	data := utils.GetData("data/2021/day8.txt")
	inputs := strings.Split(data, "\n")
	part1 := 0
	part2Values := []string{}

	for _, v := range inputs {
		parts := strings.Split(v, " | ")
		segs := strings.Split(parts[0], " ")
		keys := strings.Split(parts[1], " ")
		d := Digit{}

		sort.Slice(segs, func(i, j int) bool {
			return len(segs[i]) < len(segs[j])
		})

		mapping := make(map[string]int)

		segs2 := []string{}
		middleOrTopLeft := []string{}

		for _, v := range segs {
			if len(v) == 2 { // 1
				digits := strings.Split(v, "")
				d.Top_right = digits[0]
				d.Bottom_right = digits[1]
				mapping[sorts(v)] = 1
			} else if len(v) == 3 { // 7
				digits := strings.Split(v, "")

				for _, v := range digits {
					if v != d.Top_right && v != d.Bottom_right {
						d.Top = v
					}
				}
				mapping[sorts(v)] = 7
			} else if len(v) == 4 { // 4
				digits := strings.Split(v, "")
				for _, v := range digits {
					if v != d.Top_right && v != d.Bottom_right {
						middleOrTopLeft = append(middleOrTopLeft, v)
					}
				}
				mapping[sorts(v)] = 4
			} else if len(v) == 7 { // 8
				mapping[sorts(v)] = 8
			} else {
				segs2 = append(segs2, v) // append unfiltered
			}
		}

		segs3 := []string{}

		for _, v := range segs2 {
			digits := strings.Split(v, "")

			if len(digits) == 5 {
				if strings.Contains(v, d.Top_right) && strings.Contains(v, d.Bottom_right) { // 3
					mapping[sorts(v)] = 3
				} else if strings.Contains(v, middleOrTopLeft[0]) && strings.Contains(v, middleOrTopLeft[1]) { // 5
					mapping[sorts(v)] = 5
				} else { // 2
					mapping[sorts(v)] = 2
				}
			} else {
				segs3 = append(segs3, v)
			}
		}

		for _, v := range segs3 {
			if strings.Contains(v, d.Top_left) && strings.Contains(v, d.Middle) && strings.Contains(v, d.Top_right) && strings.Contains(v, d.Bottom_right) {
				mapping[sorts(v)] = 9
			} else if strings.Contains(v, d.Top_left) && strings.Contains(v, d.Middle) {
				mapping[sorts(v)] = 6
			} else {
				mapping[sorts(v)] = 0
			}
		}

		tempValue := ""
		for _, k := range keys {
			if len(k) == 2 || len(k) == 4 || len(k) == 3 || len(k) == 7 {
				part1 += 1
			}

			t := sorts(k)

			tempValue = fmt.Sprintf("%s%d", tempValue, mapping[t])
		}

		part2Values = append(part2Values, tempValue)
	}

	part2 := 0

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

func Translate(s string, key map[string]string) string {
	temp := ""

	for _, v := range strings.Split(s, "") {
		if _, ok := key[v]; ok {
			temp += key[v]
		}
	}

	return temp
}

func sorts(s string) string {
	l := strings.Split(s, "")
	sort.Strings(l)
	return strings.Join(l, "")
}
