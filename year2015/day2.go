package year2015

import (
	"advent-of-code/pkg/data"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Day2 -- I Was Told There Would Be No Math
func Day2() {
	data := data.ReadAsString("data/2015/day2.txt")
	boxes := strings.Split(data, "\n")

	totalArea := 0
	totalRibbon := 0

	for _, b := range boxes {
		// Part 1
		l, w, h := parse(b)
		area, small := getCubeArea(l, w, h)

		totalArea += area
		totalArea += small

		// Part 2
		ribbon := getBoxRibbonLength([]int{l, w, h})
		bowLength := getCubeVolume(l, w, h)

		totalRibbon += ribbon
		totalRibbon += bowLength
	}

	fmt.Printf("Day 2: { 1: %d, 2: %d }\n", totalArea, totalRibbon)
}

func parse(dimensions string) (int, int, int) {
	parts := strings.Split(dimensions, "x")

	l, _ := strconv.Atoi(parts[0])
	w, _ := strconv.Atoi(parts[1])
	h, _ := strconv.Atoi(parts[2])

	return l, w, h
}

// 2*l*w + 2*w*h + 2*h*l
func getCubeArea(l int, w int, h int) (int, int) {
	area := (2 * l * w) + (2 * w * h) + (2 * h * l)

	nums := []int{l, w, h}
	sort.Ints(nums)

	smallest := nums[0] * nums[1]

	return area, smallest
}

func getBoxRibbonLength(dimensions []int) int {
	sort.Ints(dimensions)
	return (dimensions[0] * 2) + (dimensions[1] * 2)
}

func getCubeVolume(l int, w int, h int) int {
	return l * w * h
}
