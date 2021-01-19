package year2015

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

type coords struct {
	aStart, aEnd int
	bStart, bEnd int
}

// Day6 -- Probably a Fire Hazard
func Day6() {
	data := utils.GetData("data/2015/day6.txt")

	list := strings.Split(data, "\n")

	coords := make([]bool, 1000000)

	for _, c := range list {
		vals := parseIns(c)

		toggleLights(&coords, vals)
	}

	fmt.Printf("Day  6: { 1: %d, 2: %d }\n", 0, len(list))
}

func parseIns(ins string) *coords {
	var data string
	var vals []string

	if strings.Contains(ins, "turn on") {
		data = strings.TrimPrefix(ins, "turn on ")
	} else if strings.Contains(ins, "turn off") {
		data = strings.TrimPrefix(ins, "turn on ")
	} else {
		data = strings.TrimPrefix(ins, "turn on ")
	}

	vals = strings.Split(data, " through ")
	a := strings.Split(vals[0], ",")
	b := strings.Split(vals[1], ",")

	aStart, _ := strconv.Atoi(a[0])
	aEnd, _ := strconv.Atoi(a[1])
	bStart, _ := strconv.Atoi(b[0])
	bEnd, _ := strconv.Atoi(b[1])

	val := coords{
		aStart,
		aEnd,
		bStart,
		bEnd,
	}

	return &val
}

func toggleLights(lights *[]bool, vals coords) {
	// To calculate where to Start & end:
	// (Y * rowSize) + X
	//
	//
}
