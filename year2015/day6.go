package year2015

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

// Coords is a coords struct
type Coords struct {
	aStart, aEnd int
	bStart, bEnd int
}

// Day6 -- Probably a Fire Hazard
func Day6() {
	data := utils.GetData("data/2015/day6.txt")
	list := strings.Split(data, "\r\n")
	coords := make([]uint8, 1000000)
	rowSize := 1000

	for _, c := range list {
		vals, action := parseIns(c)

		toggleLights(&coords, vals, action, rowSize)
		// fmt.Println(getLightCount(coords), action, vals)
	}

	litCount := getLightCount(coords)

	fmt.Printf("Day  6: { 1: %d, 2: %d }\n", litCount, len(list))
}

func getLightCount(lights []uint8) int {
	litCount := 0

	for _, c := range lights {
		if c == 1 {
			litCount++
		}
	}

	return litCount
}

func parseIns(ins string) (*Coords, string) {
	var data string
	var vals []string
	var action string

	if strings.Contains(ins, "turn on") {
		data = strings.TrimPrefix(ins, "turn on ")
		action = "TURNON"
	} else if strings.Contains(ins, "turn off") {
		data = strings.TrimPrefix(ins, "turn off ")
		action = "TURNOFF"
	} else {
		data = strings.TrimPrefix(ins, "toggle ")
		action = "TOGGLE"
	}

	vals = strings.Split(data, " through ")
	a := strings.Split(vals[0], ",")
	b := strings.Split(vals[1], ",")

	aStart, _ := strconv.Atoi(a[0])
	aEnd, _ := strconv.Atoi(a[1])
	bStart, _ := strconv.Atoi(b[0])
	bEnd, _ := strconv.Atoi(b[1])

	val := Coords{
		aStart,
		aEnd,
		bStart,
		bEnd,
	}

	return &val, action
}

func toggleLights(lights *[]uint8, vals *Coords, action string, rowSize int) {
	// To calculate where to Start & end:
	// (Y * rowSize) + X

	start := (vals.aEnd * rowSize) + vals.aStart
	end := (vals.bEnd * rowSize) + vals.bStart

	// fmt.Printf("%s - s: {%d, %d} e: {%d, %d} Count: %d \n", action, vals.aStart, vals.aEnd, vals.bStart, vals.bEnd, (end - start))

	// Potentially need end + 1
	for i := start; i <= end; i++ {
		switch action {
		case "TURNON":
			(*lights)[i] = 1
		case "TURNOFF":
			(*lights)[i] = 0
		case "TOGGLE":
			if (*lights)[i] == 1 {
				(*lights)[i] = 0
			} else {
				(*lights)[i] = 1
			}
		}
	}
}

func fancyPrint(lights *[]uint8, rowSize int) {
	for i := 0; i < rowSize; i++ {
		row := (*lights)[i:((rowSize * i) + rowSize)]
		rowString := strings.Trim(strings.Join(strings.Split(fmt.Sprint(row), " "), ""), "[]")
		fmt.Println(rowString)
	}
}
