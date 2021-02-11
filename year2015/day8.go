package year2015

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
)

// Day6 -- Probably a Fire Hazard
func Day8() {
	data := utils.GetDataBytes("data/2015/day8.txt")
	// data := utils.GetDataBytes("data/2015/day8_test.txt")

	real := 0
	mem := 0
	newEncoding := 0

	var words []string
	currentWord := ""

	for i, c := range data {
		if c == '\n' || i == (len(data)-1) {
			if c != '\n' {
				if c == 0x22 || c == 0x5c {
					newEncoding += 2
				} else {
					newEncoding++
				}

				real++
				currentWord = currentWord + string(c)
			}

			words = append(words, currentWord)

			fmt.Printf("Current Count for %s: R -> %d, New -> %d\n", currentWord, len(currentWord), newEncoding+(len(words)*2))
			memRep, err := strconv.Unquote(currentWord)

			if err != nil {
				fmt.Println(err)
			}

			mem += len(memRep)
			currentWord = ""
		} else {
			fmt.Printf("%x - %c\n", c, c)

			if c == 0x22 || c == 0x5c {
				newEncoding += 2
			} else {
				newEncoding++
			}

			real++
			currentWord = currentWord + string(c)
		}
	}

	newEncoding = newEncoding + (len(words) * 2)
	// fmt.Printf("Real: %d Mem: %d New: %d\n", real, mem, newEncoding)

	fmt.Printf("Day 8: { 1: %d, 2: %d }\n", (real - mem), (newEncoding - real))
}
