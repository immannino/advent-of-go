package year2021

import (
	"advent-of-code/pkg/data"
	"fmt"
	"strconv"
	"strings"
)

// --- Day 3: Binary Diagnostic ---
func Day3() {
	data := data.ReadAsString("data/2021/day3.txt")
	rows := strings.Split(data, "\n")
	gammaStr := ""   // Most common bit
	epsilonStr := "" // Inverse of Gamma
	counts := make([]int, len(rows[0]))

	for _, v := range rows {
		bits := strings.Split(v, "")

		for i, b := range bits {
			val, err := strconv.Atoi(b)
			if err != nil {
				panic(err)
			}
			counts[i] += val
		}
	}

	for _, v := range counts {
		if v >= (len(rows) / 2) {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaStr, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)
	part1 := gamma * epsilon

	oxyStr := filter(rows, "", true, 0, len(rows[0]))
	co2Str := filter(rows, "", false, 0, len(rows[0]))

	oxy, _ := strconv.ParseInt(oxyStr, 2, 64)
	co2, _ := strconv.ParseInt(co2Str, 2, 64)
	part2 := co2 * oxy

	fmt.Printf("--- Day 3: Binary Diagnostic --- { 1: %d, 2: %d }\n", part1, part2)
}

func filter(rows []string, str string, common bool, col int, rowSize int) string {
	nextRows := []string{}
	nextStr := str
	commonChar := ""
	charCount := 0
	size := len(rows)

	if size == 0 || col == rowSize {
		return str
	}

	if size == 1 && col < rowSize {
		return rows[0]
	}

	for _, v := range rows {
		val, _ := strconv.Atoi(string(v[col]))
		charCount += val
	}

	if charCount > (size/2) || charCount == (size/2) && size%2 == 0 {
		commonChar = "1"
	} else {
		commonChar = "0"
	}

	if commonChar == "1" {
		if common {
			nextStr += "1"
		} else {
			nextStr += "0"
		}
	} else {
		if common {
			nextStr += "0"
		} else {
			nextStr += "1"
		}
	}

	for _, v := range rows {
		if common {
			if commonChar == string(v[col]) {
				nextRows = append(nextRows, v)
			}
		} else {
			if commonChar != string(v[col]) {
				nextRows = append(nextRows, v)
			}
		}
	}

	return filter(nextRows, nextStr, common, col+1, rowSize)
}
