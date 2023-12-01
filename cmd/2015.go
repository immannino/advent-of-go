package cmd

import (
	"advent-of-code/internal"
	"advent-of-code/internal/data"
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func NewYear2015() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1: Not Quite Lisp", Day1_2015),
		internal.NewPuzzle(2, "Day 2: I Was Told There Would Be No Math", Day2_2015),
		internal.NewPuzzle(3, "Day 3: Perfectly Spherical Houses in a Vacuum", Day3_2015),
		internal.NewPuzzle(4, "Day 4: Perfectly Spherical Houses in a Vacuum", Day4_2015),
		internal.NewPuzzle(5, "Day 5: Doesn't He Have Intern-Elves For This?", Day5_2015),
		internal.NewPuzzle(6, "Day 6: Probably a Fire Hazard", Day6_2015),
		internal.NewPuzzle(8, "Day 8: Matchsticks", Day8_2015),
		internal.NewPuzzle(8, "Day 9: All in a Single Night", Day9_2015),
	}

	return internal.Year{"Year 2015", Days}
}

// Day1 -- Not Quite Lisp
func Day1_2015() internal.Answer {
	data := data.ReadAsString("data/2015/day1.txt")

	floor := 0
	input := []rune(data)
	position := 1
	hasMatch := false

	for _, e := range input {
		if e == '(' {
			floor += 1
		} else {
			floor -= 1
		}

		if !hasMatch && floor == -1 {
			hasMatch = true
		}

		if !hasMatch {
			position += 1
		}
	}

	return internal.Answer{
		Part1: strconv.Itoa(floor),
		Part2: strconv.Itoa(position),
	}
}

func Day2_2015() internal.Answer {
	data := data.ReadAsString("data/2015/day2.txt")
	boxes := strings.Split(data, "\n")

	totalArea := 0
	totalRibbon := 0

	for _, b := range boxes {
		// Part 1
		l, w, h := parseDay2(b)
		area, small := getCubeArea(l, w, h)

		totalArea += area
		totalArea += small

		// Part 2
		ribbon := getBoxRibbonLength([]int{l, w, h})
		bowLength := getCubeVolume(l, w, h)

		totalRibbon += ribbon
		totalRibbon += bowLength
	}

	return internal.Answer{
		Part1: strconv.Itoa(totalArea),
		Part2: strconv.Itoa(totalRibbon),
	}
}

func parseDay2(dimensions string) (int, int, int) {
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

func Day3_2015() internal.Answer {
	data := data.ReadAsString("data/2015/day3.txt")
	coords := []rune(data)
	visited := make(map[string]bool)
	pos := []int{0, 0}

	visited["{0,0}"] = true

	// Part 1
	for _, c := range coords {
		switch c {
		case '^':
			pos[1] += 1
		case '>':
			pos[0] += 1
		case 'v':
			pos[1] -= 1
		case '<':
			pos[0] -= 1
		}

		coordString := fmt.Sprintf("{%d,%d}", pos[0], pos[1])

		if !visited[coordString] {
			visited[coordString] = true
		}
	}

	// Part 2
	roboVisited := make(map[string]bool)
	robo := []int{0, 0}
	santa := []int{0, 0}

	roboVisited["{0,0}"] = true

	for i := 0; i < len(coords); i += 2 {
		santaAction := coords[i]
		roboAction := coords[i+1]

		switch santaAction {
		case '^':
			santa[1] += 1
		case '>':
			santa[0] += 1
		case 'v':
			santa[1] -= 1
		case '<':
			santa[0] -= 1
		}

		coordString := fmt.Sprintf("{%d,%d}", santa[0], santa[1])

		if !roboVisited[coordString] {
			roboVisited[coordString] = true
		}

		switch roboAction {
		case '^':
			robo[1] += 1
		case '>':
			robo[0] += 1
		case 'v':
			robo[1] -= 1
		case '<':
			robo[0] -= 1
		}

		coordString = fmt.Sprintf("{%d,%d}", robo[0], robo[1])

		if !roboVisited[coordString] {
			roboVisited[coordString] = true
		}
	}

	return internal.Answer{
		Part1: strconv.Itoa(len(visited)),
		Part2: strconv.Itoa(len(roboVisited)),
	}
}

func Day4_2015() internal.Answer {
	// data := "yzbqklnj"

	h := md5.New()
	// secret := []byte(data4)

	io.WriteString(h, "00000")

	// fmt.Printf("%x", h.Sum(secret))

	// fmt.Printf("Day 4: { 1: %d, 2: %d }\n", h.Sum(nil), 0)
	// fmt.Printf("Day 4: Learn MD5 hash\n")
	return internal.Answer{}
}

func Day5_2015() internal.Answer {
	data := data.ReadAsString("data/2015/day5.txt")
	list := strings.Split(data, "\n")
	count := 0
	count2 := 0

	for _, c := range list {
		// part 1
		if hasThreeVowel(c) && hasRepeatingChar(c) && doesNotHaveSpecificStrings(c) {
			count++
		}

		if hasOverlappingPairs(c) && ovo(c) {
			count2++
		}
	}

	return internal.Answer{
		Part1: strconv.Itoa(count),
		Part2: strconv.Itoa(count2),
	}

}

// ---- Part 1

// Checks for the presence of at least 3 vowels
// Supported: 'aeiou'
func hasThreeVowel(word string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	count := 0

	for _, c := range word {
		if contains(&vowels, c) {
			count++
		}
	}

	if count >= 3 {
		return true
	}

	return false
}

func contains(valid *[]rune, letter rune) bool {
	for _, c := range *valid {
		if letter == c {
			return true
		}
	}

	return false
}

func hasRepeatingChar(word string) bool {
	currentChar := rune(word[0])

	for i := 1; i < len(word); i++ {
		tempChar := rune(word[i])

		if tempChar == currentChar {
			return true
		}

		currentChar = tempChar
	}

	return false
}

// Checks for the absence of specific strings:
// Strings: [ ab, cd, pq, or xy ]
func doesNotHaveSpecificStrings(word string) bool {
	current := rune(word[0])
	badVals := []string{"ab", "cd", "pq", "xy"}

	for i := 1; i < len(word); i++ {
		temp := rune(word[i])

		str := string([]rune{current, temp})

		if isBadString(&badVals, str) {
			return false
		}

		current = temp
	}

	return true
}

func isBadString(bad *[]string, word string) bool {
	for _, c := range *bad {
		if c == word {
			return true
		}
	}

	return false
}

// -- Part 2

// Create a map that keeps count of pairs. Return true if any keys have count 2+
func hasOverlappingPairs(word string) bool {
	for i := 0; i < (len(word) - 1); i++ {
		substr := word[i:(i + 2)]

		parts := strings.Split(word, substr)

		if len(parts) >= 3 {
			return true
		}
	}

	return false
}

// Returns true if any 3 chars follow the OvO pattern
func ovo(word string) bool {
	for i := 0; i < (len(word) - 2); i++ {
		first := word[i]
		last := word[i+2]

		if first == last {
			return true
		}
	}

	return false
}

// Coords is a coords struct
type Coords struct {
	aStart, aEnd int
	bStart, bEnd int
}

func Day6_2015() internal.Answer {
	data := data.ReadAsString("data/2015/day6.txt")
	list := strings.Split(data, "\r\n")
	coords := make([]uint8, 1000000)
	coords2 := make([]uint8, 1000000)
	rowSize := 1000

	for _, c := range list {
		vals, action := parseIns(c)

		toggleLights(&coords, vals, action, rowSize)
		toggleLightsPart2(&coords2, vals, action, rowSize)
	}

	litCount := getLightCount(coords)
	litCount2 := getLightPower(coords2)

	return internal.Answer{
		Part1: strconv.Itoa(litCount),
		Part2: strconv.Itoa(litCount2),
	}
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

func getLightPower(lights []uint8) int {
	power := 0

	for _, c := range lights {
		power += int(c)
	}

	return power
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

	for x := vals.aStart; x <= vals.bStart; x++ {
		for y := vals.aEnd; y <= vals.bEnd; y++ {
			i := (rowSize * x) + y

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
}

func toggleLightsPart2(lights *[]uint8, vals *Coords, action string, rowSize int) {

	for x := vals.aStart; x <= vals.bStart; x++ {
		for y := vals.aEnd; y <= vals.bEnd; y++ {
			i := (rowSize * x) + y

			switch action {
			case "TURNON":
				(*lights)[i]++
			case "TURNOFF":
				if (*lights)[i] > 0 {
					(*lights)[i]--
				}
			case "TOGGLE":
				(*lights)[i] += 2
			}
		}
	}
}

func fancyPrint(lights *[]uint8, rowSize int) {
	// for i := 0; i < rowSize; i++ {
	// 	row := (*lights)[i:((rowSize * i) + rowSize)]
	// 	rowString := strings.Trim(strings.Join(strings.Split(fmt.Sprint(row), " "), ""), "[]")
	// 	fmt.Println(rowString)
	// }
}

func Day8_2015() internal.Answer {
	data := data.ReadAsBytes("data/2015/day8.txt")

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

			// fmt.Printf("Current Count for %s: R -> %d, New -> %d\n", currentWord, len(currentWord), newEncoding+(len(words)*2))
			memRep, err := strconv.Unquote(currentWord)

			if err != nil {
				fmt.Println(err)
			}

			mem += len(memRep)
			currentWord = ""
		} else {
			// fmt.Printf("%x - %c\n", c, c)

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

	return internal.Answer{
		Part1: strconv.Itoa((real - mem)),
		Part2: strconv.Itoa((newEncoding - real)),
	}
}

func Day9_2015() internal.Answer {
	// input := data.ReadAsString("data/2015/day9.txt")
	return internal.Answer{}
}
