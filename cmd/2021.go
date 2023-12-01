package cmd

import (
	"advent-of-code/internal"
	"advent-of-code/internal/data"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func NewYear2021() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1: Sonar Sweep", Day1_2021),
		internal.NewPuzzle(2, "Day 2: Dive!", Day2_2021),
		internal.NewPuzzle(3, "Day 3: Binary Diagnostic", Day3_2021),
		internal.NewPuzzle(4, "Day 4: Giant Squid", Day4_2021),
		internal.NewPuzzle(5, "Day 5: Hydrothermal Venture", Day5_2021),
		internal.NewPuzzle(6, "Day 6: Lanternfish", Day6_2021),
		internal.NewPuzzle(7, "Day 7: The Treachery of Whales", Day7_2021),
		internal.NewPuzzle(8, "Day 8: Seven Segment Search", Day8_2021),
	}

	return internal.Year{"Year 2021", Days}
}

func Day1_2021() internal.Answer {
	data := data.ReadAsString("data/2021/day1.txt")
	rows := strings.Split(data, "\n")
	part1 := -1
	previous := 0

	// part 1
	for _, v := range rows {
		i, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		if i > previous {
			part1 += 1
		}

		previous = i
	}

	part2 := -1
	previous = 0

	for i := 0; i < len(rows)-2; i++ {
		m := getTMSWSum(i, rows)

		if m > previous {
			part2 += 1
		}

		previous = m
	}

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func getTMSWSum(start int, rows []string) int {
	sum := 0

	for i := start; i < start+3; i++ {
		num, err := strconv.Atoi(rows[i])

		if err != nil {
			panic(err)
		}

		sum += num
	}

	return sum
}

func Day2_2021() internal.Answer {
	data := data.ReadAsString("data/2021/day2.txt")
	rows := strings.Split(data, "\n")
	part1 := len(data)
	part2 := 0

	depth := 0
	depth2 := 0
	horizontal := 0
	aim := 0

	for _, v := range rows {
		if strings.Contains(v, "forward") {
			val := getNum(v, "forward ")
			horizontal += val
			depth2 += val * aim
		} else if strings.Contains(v, "up") {
			val := getNum(v, "up ")
			depth -= val
			aim -= val
		} else if strings.Contains(v, "down") {
			val := getNum(v, "down ")
			depth += val
			aim += val
		}
	}

	part1 = horizontal * depth
	part2 = horizontal * depth2

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func getNum(input string, sep string) int {
	numString := strings.Split(input, sep)[1]
	num, err := strconv.Atoi(numString)

	if err != nil {
		panic(err)
	}

	return num
}

func Day3_2021() internal.Answer {
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

	return internal.Answer{
		Part1: strconv.Itoa(int(part1)),
		Part2: strconv.Itoa(int(part2)),
	}
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

type Board struct {
	hasWon bool
	nums   [][]BingoNumber
}

type BingoNumber struct {
	number  int
	checked bool
}

func Day4_2021() internal.Answer {
	data := data.ReadAsString("data/2021/day4.txt")
	rows := strings.Split(data, "\n\n")
	bingoNumStrings, rows := strings.Split(rows[0], ","), rows[1:]
	bingoNums := make([]int, len(bingoNumStrings))

	for k, v := range bingoNumStrings {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		bingoNums[k] = n
	}

	boards := []Board{}

	for _, v := range rows {
		tempBoard := Board{}
		tempRows := strings.Split(v, "\n")

		for _, r := range tempRows {
			rowNums := strings.Split(r, " ")
			tempRow := []BingoNumber{}
			for _, num := range rowNums {
				if len(num) > 0 {
					val, err := strconv.Atoi(num)
					if err != nil {
						panic(err)
					}

					tempRow = append(tempRow, BingoNumber{number: val, checked: false})
				}
			}

			tempBoard.nums = append(tempBoard.nums, tempRow)
		}
		boards = append(boards, tempBoard)
	}

	activeBingoNumber := bingoNums[0]
	activeBingoNumIndex := 0
	hasWinner := false
	winnerIndex := -1

	for !hasWinner {
		for index, _ := range boards {
			boards[index].check(activeBingoNumber)
		}

		hasWinner, winnerIndex = checkWinners(boards)

		if !hasWinner {
			activeBingoNumIndex += 1
			activeBingoNumber = bingoNums[activeBingoNumIndex]
		}
	}

	winnerUncheckedSum := boards[winnerIndex].sumUnchecked()
	part1 := winnerUncheckedSum * activeBingoNumber

	for len(checkNonWinners(boards)) > 1 {
		for index, _ := range boards {
			boards[index].check(activeBingoNumber)
		}

		if len(checkNonWinners(boards)) > 1 {
			activeBingoNumIndex += 1
			activeBingoNumber = bingoNums[activeBingoNumIndex]
		}
	}

	lastWinners := checkNonWinners(boards)
	lastWinner := lastWinners[0]

	activeBingoNumber = bingoNums[0]
	activeBingoNumIndex = 0

	for lastWinner.hasWon == false {
		lastWinner.check(activeBingoNumber)

		if lastWinner.hasWon == false {
			activeBingoNumIndex += 1
			activeBingoNumber = bingoNums[activeBingoNumIndex]
		}
	}

	unchecked := lastWinner.sumUnchecked()

	part2 := unchecked * activeBingoNumber

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func prettyPrint(b []Board) {
	for _, board := range b {
		temp := ""

		for _, r := range board.nums {
			for _, c := range r {
				if c.checked {
					temp += "1 "
				} else {
					temp += "0 "
				}
			}
			temp += "\n"
		}

		// fmt.Println(temp + "\n")
	}
}

func (b *Board) check(num int) {
	for rowIndex, row := range b.nums {
		for colIndex, col := range row {
			if col.number == num {
				b.nums[rowIndex][colIndex].checked = true
			}
		}
	}

	b.hasWon = b.checkHasWon()
}

func (b *Board) checkHasWon() bool {
	// Check Cols
	for _, r := range b.nums {
		isMatch := true
		for _, c := range r {
			isMatch = isMatch && c.checked
		}

		if isMatch {
			return true
		}
	}

	// Check Rows
	for col := 0; col < len(b.nums); col++ {
		isMatch := true

		for row := 0; row < len(b.nums); row++ {
			isMatch = isMatch && b.nums[row][col].checked
		}

		if isMatch {
			return true
		}

	}

	return false
}

func (b *Board) sumUnchecked() int {
	sum := 0

	for _, row := range b.nums {
		for _, col := range row {
			if !col.checked {
				sum += col.number
			}
		}
	}

	return sum
}

func checkWinners(b []Board) (bool, int) {
	for i, v := range b {
		if v.hasWon {
			return true, i
		}
	}

	return false, -1
}

func checkNonWinners(b []Board) []Board {
	temp := []Board{}

	for _, v := range b {
		if !v.hasWon {
			temp = append(temp, v)
		}
	}

	return temp
}

func Day5_2021() internal.Answer {
	data := data.ReadAsString("data/2021/day5.txt")
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

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
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

func Day6_2021() internal.Answer {
	data := data.ReadAsString("data/2021/day6.txt")
	fishes := make([]int, 9)

	for _, v := range strings.Split(data, ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		fishes[n] += 1
	}

	part1 := 0
	for day := 0; day < 256; day++ {
		tempFishes := make([]int, 9)

		if day == 80 {
			for k := range fishes {
				part1 += fishes[k]
			}
		}

		for k, v := range fishes {
			if k == 0 {
				tempFishes[6] += v
				tempFishes[8] += v
			} else {
				tempFishes[k-1] += v
			}
		}

		copy(fishes, tempFishes)
	}

	part2 := 0
	for _, v := range fishes {
		part2 += v
	}

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func Day7_2021() internal.Answer {
	data := data.ReadAsString("data/2021/day7.txt")
	rows := strings.Split(data, ",")
	crabs, min, max := getCrabs(rows)
	horizontals := make([]int, (min+max)+1)
	horizontalsMultiplied := make([]int, (min+max)+1)

	for i := min; i <= max; i++ {
		diff := 0
		diff2 := 0

		for _, v := range crabs {
			if v > i {
				diff += (v - i)
				diff2 += summation((v - i))
			} else {
				diff += (i - v)
				diff2 += summation((i - v))
			}
		}

		horizontals[i] = diff
		horizontalsMultiplied[i] = diff2
	}

	smallest := 0
	smallestMult := 0

	for i := 0; i < len(horizontals); i++ {
		if horizontals[i] < horizontals[smallest] {
			smallest = i
		}

		if horizontalsMultiplied[i] < horizontalsMultiplied[smallestMult] {
			smallestMult = i
		}
	}

	part1 := horizontals[smallest]
	part2 := horizontalsMultiplied[smallestMult]

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func summation(n int) int {
	if n%2 == 0 {
		return (n / 2) * (n + 1)
	}

	return ((n + 1) / 2) * n
}

func getCrabs(rows []string) ([]int, int, int) {
	nums := []int{}

	n, err := strconv.Atoi(rows[0])
	if err != nil {
		panic(err)
	}

	nums = append(nums, n)
	min := n
	max := n

	for i := 1; i < len(rows); i++ {
		n, err := strconv.Atoi(rows[i])
		if err != nil {
			panic(err)
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}

		nums = append(nums, n)
	}

	return nums, min, max
}

func Day8_2021() internal.Answer {
	data := data.ReadAsString("data/2021/day8.txt")
	inputs := strings.Split(data, "\n")
	part1 := 0
	part2 := 0
	part2Values := []string{}
	sumMap := map[int]int{
		42: 0,
		17: 1,
		34: 2,
		39: 3,
		30: 4,
		37: 5,
		41: 6,
		25: 7,
		49: 8,
		45: 9,
	}

	for _, v := range inputs {
		parts := strings.Split(v, " | ")
		segs := strings.Split(parts[0], " ")
		keys := strings.Split(parts[1], " ")

		sort.Slice(segs, func(i, j int) bool {
			return len(segs[i]) < len(segs[j])
		})

		charMap := make(map[string]int)

		// Score the weights
		for _, v := range segs {
			digits := strings.Split(v, "")

			for _, digit := range digits {
				if _, ok := charMap[digit]; ok {
					charMap[digit] += 1
				} else {
					charMap[digit] = 1
				}
			}
		}

		tempNum := ""

		for _, k := range keys {
			if len(k) == 2 || len(k) == 4 || len(k) == 3 || len(k) == 7 {
				part1 += 1
			}

			digits := strings.Split(k, "")
			keySum := 0

			for _, d := range digits {
				keySum += charMap[d]
			}

			tempNum = fmt.Sprintf("%s%d", tempNum, sumMap[keySum])
		}

		n, err := strconv.Atoi(tempNum)
		if err != nil {
			panic(err)
		}

		part2 += n
	}

	for _, v := range part2Values {
		// fmt.Println(v)
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		part2 += n
	}

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func sorts(s string) string {
	l := strings.Split(s, "")
	sort.Strings(l)
	return strings.Join(l, "")
}
