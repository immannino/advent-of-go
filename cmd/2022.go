package cmd

import (
	"advent-of-code/internal"
	"advent-of-code/internal/data"
	"advent-of-code/internal/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func NewYear2022() internal.Year {
	Days := []internal.PuzzleInterface{
		internal.NewPuzzle(1, "Day 1: Calorie Counting", Day1_2022),
		internal.NewPuzzle(2, "Day 2: Rock Paper Scissors", Day2_2022),
		internal.NewPuzzle(3, "Day 3: Rocksack Reorganization", Day3_2022),
		internal.NewPuzzle(4, "Day 4: Camp Cleanup", Day4_2022),
		internal.NewPuzzle(5, "Day 5: Supply Stacks", Day5_2022),
		internal.NewPuzzle(6, "Day 6: Tuning Trouble", Day6_2022),
		internal.NewPuzzle(7, "Day 7: No Space Left on Device", Day7_2022),
	}

	return internal.Year{"Year 2022", Days}
}

func Day1_2022() internal.Answer {
	var input string
	var part1 int
	var part2 int
	// var example string = `1000
	// 2000
	// 3000

	// 4000

	// 5000
	// 6000

	// 7000
	// 8000
	// 9000

	// 10000
	// `
	input = data.ReadAsString("data/2022/day1.txt")
	// data = example
	rows := strings.Split(input, "\n")

	elves := []int{}
	currentElf := 0

	for _, r := range rows {
		if r == "" {
			elves = append(elves, currentElf)
			currentElf = 0
			continue
		}

		v, _ := strconv.ParseInt(r, 10, 32)

		currentElf += int(v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	part1 = elves[0]

	part2 = elves[0] + elves[1] + elves[2]

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

var oppHands = map[string]string{
	"A": "ROCK",
	"B": "PAPER",
	"C": "SCISSORS",
}

var playerHands = map[string]string{
	"X": "ROCK",
	"Y": "PAPER",
	"Z": "SCISSORS",
}

var roundFinishes = map[string]string{
	"X": "LOSE",
	"Y": "DRAW",
	"Z": "WIN",
}

var handScores = map[string]int{
	"ROCK":     1,
	"PAPER":    2,
	"SCISSORS": 3,
}

var outcomes = map[string]int{
	"WIN":  6,
	"DRAW": 3,
	"LOSE": 0,
}

var playerOppWinOutcomes = map[string]bool{
	"ROCK-PAPER":     false,
	"ROCK-SCISSORS":  true,
	"PAPER-ROCK":     true,
	"PAPER-SCISSORS": false,
	"SCISSORS-ROCK":  false,
	"SCISSORS-PAPER": true,
}

func Day2_2022() internal.Answer {
	var input string
	var part1 int
	var part2 int
	// 	var example string = `A Y
	// B X
	// C Z`
	input = data.ReadAsString("data/2022/day2.txt")
	rounds := strings.Split(input, "\n")

	for _, r := range rounds {
		hands := strings.Split(r, " ")

		// Part 1
		outcome := judge(playerHands[hands[1]], oppHands[hands[0]])
		score := outcomes[outcome] + handScores[playerHands[hands[1]]]
		part1 += score

		// Part 2
		player := determinePlayerMove(roundFinishes[hands[1]], oppHands[hands[0]])
		newScore := outcomes[roundFinishes[hands[1]]] + handScores[player]
		part2 += newScore
	}

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func judge(player string, opp string) string {
	if opp == player {
		return "DRAW"
	}

	outcome := playerOppWinOutcomes[fmt.Sprintf("%s-%s", player, opp)]

	if outcome {
		return "WIN"
	}

	return "LOSE"
}

func determinePlayerMove(outcome string, opp string) string {
	if outcome == "DRAW" {
		return opp
	}

	if outcome == "LOSE" {
		switch opp {
		case "ROCK":
			return "SCISSORS"
		case "PAPER":
			return "ROCK"
		case "SCISSORS":
			return "PAPER"
		}
	}

	if outcome == "WIN" {
		switch opp {
		case "ROCK":
			return "PAPER"
		case "PAPER":
			return "SCISSORS"
		case "SCISSORS":
			return "ROCK"
		}
	}

	return ""
}

func Day3_2022() internal.Answer {
	var part1 int
	var part2 int

	// Index of letter is priority value numerical - e.g. alphabet['k'] == index + 1 priority
	var alphabet string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`

	// 	var example string = `vJrwpWtwJgWrhcsFMMfFFhFp
	// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	// PmmdzqPrVvPwwTWBwg
	// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	// ttgJtRGJQctTZtZT
	// CrZsJsPPZsGzwwsLwLmpwMDw
	// `
	input := data.ReadAsString("data/2022/day3.txt")
	compartments := strings.Split(input, "\n")

	// Part 1
	for _, c := range compartments {
		left := c[:len(c)/2] // left half of the compartment to compare against
		right := c[len(c)/2:]

		for _, p := range right {
			if strings.Contains(left, string(p)) {
				part1 += strings.Index(alphabet, string(p)) + 1
				break
			}
		}
	}

	// Part 2 - Segments of 3
	for i := 0; i < len(compartments)/3; i++ {
		segs := compartments[i*3 : (i*3)+3] // Groups of 3 inputs

		for _, p := range segs[0] {
			if strings.Contains(segs[1], string(p)) && strings.Contains(segs[2], string(p)) {
				part2 += strings.Index(alphabet, string(p)) + 1
				break
			}
		}
	}

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func Day4_2022() internal.Answer {
	var part1 int
	var part2 int
	// 	var example string = `2-4,6-8
	// 2-3,4-5
	// 5-7,7-9
	// 2-8,3-7
	// 6-6,4-6
	// 2-6,4-8`
	input := data.ReadAsString("data/2022/day4.txt")
	coords := strings.Split(input, "\n")

	for _, c := range coords {
		// time.Sleep(time.Second * 3)
		parts := strings.Split(c, ",")

		aStart, aEnd := parseDay42022(parts[0])
		bStart, bEnd := parseDay42022(parts[1])

		// Part 1 comparisons
		if bStart >= aStart && bEnd <= aEnd {
			part1 += 1
			continue
		}

		if aStart >= bStart && aEnd <= bEnd {
			part1 += 1
		}
	}

	for _, c := range coords {
		parts := strings.Split(c, ",")
		aStart, aEnd := parseDay42022(parts[0])
		bStart, bEnd := parseDay42022(parts[1])

		m := make(map[int]bool)

		for i := aStart; i <= aEnd; i++ {
			m[i] = true
		}

		for i := bStart; i <= bEnd; i++ {
			if _, exists := m[i]; exists {
				part2 += 1
				break
			}
		}
	}

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}

}

func parseDay42022(s string) (int, int) {
	parts := strings.Split(s, "-")
	a, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		panic(err)
	}

	b, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		panic(err)
	}

	if s != fmt.Sprintf("%d-%d", a, b) {
		panic("parse error")
	}

	return int(a), int(b)
}

func Day5_2022() internal.Answer {
	var part1 string
	var part2 string

	input := data.ReadAsString("data/2022/day5.txt")
	instructions := strings.Split(input, "\n\n")
	stacks, cols := parseBoxes(instructions[0])
	stacksPart2 := make([]string, cols)

	copy(stacksPart2, stacks)

	actions := strings.Split(instructions[1], "\n")

	// fmt.Println(len(stacks), len(stacksPart2), len(actions))
	for _, line := range actions {
		// Part 1
		count, from, to := parseSupplyLine(line)
		// Find items to move
		sub := stacks[from][len(stacks[from])-count:]
		// Find new state of From
		newFrom := stacks[from][:len(stacks[from])-count]

		// Set new From
		stacks[from] = newFrom
		// Append to To
		stacks[to] += util.Reverse(sub)

		// Part 2 (the same, minus util.Reverse)
		sub2 := stacksPart2[from][len(stacksPart2[from])-count:]
		newFrom2 := stacksPart2[from][:len(stacksPart2[from])-count]
		stacksPart2[from] = newFrom2
		stacksPart2[to] += string(sub2)
	}

	for _, v := range stacks {
		if len(v) > 0 {
			part1 += string(v[len(v)-1])
		}
	}

	for _, v := range stacksPart2 {
		if len(v) > 0 {
			part2 += string(v[len(v)-1])
		}
	}

	return internal.Answer{
		Part1: part1,
		Part2: part2,
	}
}

func print(s []string) {
	for i, v := range s {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func parseBoxes(input string) ([]string, int) {
	lines := strings.Split(input, "\n")
	linesWithout := lines[:len(lines)-1]
	cols := (len(lines[0]) / 4) + 1
	boxes := make([]string, cols)

	for _, l := range linesWithout {
		l += " " // Pad a space
		for i := 0; i < cols; i++ {
			item := l[i*4 : (i*4)+4]
			if string(item[0]) == "[" {
				val := strings.Trim(item, " []")
				boxes[i] += val
			}
		}
	}

	for i := range boxes {
		boxes[i] = util.Reverse(boxes[i])
	}

	return boxes, cols
}

func parseSupplyLine(i string) (int, int, int) {
	parts := strings.Split(strings.Replace(
		strings.Replace(
			strings.Replace(i, "move ", "", 1),
			" from ",
			",",
			1),
		" to ",
		",",
		1), ",")

	count, err := strconv.ParseInt(parts[0], 10, 16)
	if err != nil {
		panic(err)
	}

	from, err := strconv.ParseInt(parts[1], 10, 16)
	if err != nil {
		panic(err)
	}

	to, err := strconv.ParseInt(parts[2], 10, 16)
	if err != nil {
		panic(err)
	}

	return int(count), int(from) - 1, int(to) - 1

}

func Day6_2022() internal.Answer {
	input := data.ReadAsString("data/2022/day6.txt")
	part1 := findDay62022(input, 4)
	part2 := findDay62022(input, 14)

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func findDay62022(input string, coef int) int {
	for i := range input {
		next := input[i : i+coef]
		has := make(map[rune]bool)
		fail := false

		for _, v := range next {
			if _, ok := has[v]; ok {
				fail = true
				break
			}

			has[v] = true
		}

		if !fail {
			return i + coef
		}
	}

	return 0
}

const (
	PART_1_THRESHOLD = 100000
)

type File struct {
	name     string
	fileType string
	size     int
	children []*File
	parent   *File
	depth    int
}

func ParentOf(child *File, parent *File) *File {
	// fmt.Printf("Looking for parent of %s, have %s, type: %s\n", name, parent.name, parent.fileType)
	if parent.fileType == "file" {
		return nil
	}

	for _, c := range parent.children {
		if c.name == child.name && c.depth == child.depth {
			return parent
		}
	}

	for _, c := range parent.children {
		r := ParentOf(child, c)
		if r != nil {
			return r
		}
	}

	// For compiler
	return nil
}

// DFS
func Find(name string, depth int, node *File) *File {
	// fmt.Printf("Looking for: %s, At: %s\n", name, node.name)

	if name == node.name && depth == node.depth {
		return node
	}

	for _, n := range node.children {
		// Only crawl dirs
		if n.fileType == "dir" {
			r := Find(name, depth, n)
			if r != nil {
				return r
			}
		}
	}

	// For compiler
	return nil
}

func FindBFS(name string, depth int, node *File) *File {
	queue := []*File{node}

	for len(queue) > 0 {
		current := queue[0] // Fetch
		queue = queue[1:]   // Pop

		if current.name == name && current.depth == depth {
			return current
		}

		queue = append(queue, current.children...)
	}

	return nil
}

func Sum(sum uint64, node *File) uint64 {
	// fmt.Println(sum, node.name)
	if node.fileType == "file" {
		return uint64(node.size)
	}

	if node.fileType == "dir" && len(node.children) == 0 {
		return 0
	}

	for _, c := range node.children {
		sum += Sum(sum, c)
	}

	return sum
}

func Walk(depth int, node *File) *File {
	// Pint stuf f
	return nil
}

func Day7_2022() internal.Answer {
	input := data.ReadAsString("data/2022/day7.txt")
	part1 := 0
	part2 := 0
	cmds := strings.Split(input, "\n")
	root := &File{
		name:     "root",
		fileType: "dir",
		children: []*File{},
		depth:    0,
		parent:   nil,
	}

	// Build tree
	var curr *File

	for _, line := range cmds {
		if curr != nil {
			// fmt.Println("command: ", line, curr.name, len(curr.children))

			// time.Sleep(1 * time.Second)
			// fmt.Println()
			if curr.name != "root" {
				// PrintDir(ParentOf(curr, root))
				// PrintDir(root)
			}
			// fmt.Println()
		}
		if strings.Contains(line, "$ cd") {
			path := strings.Replace(line, "$ cd ", "", 1)

			if path == ".." {
				// traverse up
				if curr.name != "root" {
					// fmt.Println("Searching for Parent of: ", curr.name)
					// curr = ParentOf(curr, root)
					curr = curr.parent
					// fmt.Println("Found Parent: ", curr.name)
				}
				// if curr.name != "root" {
				// fmt.Println("Searching for Parent of: ", curr.name)
				// curr = ParentOf(curr.name, root)
				// fmt.Println("Found Parent: ", curr.name)
				// }
			} else if path == "/" {
				// fmt.Println("Set root")
				curr = root
			} else {
				// Switch to previously declared node
				curr = FindBFS(path, curr.depth+1, root)
				if path != curr.name {
					panic(fmt.Sprintf("Found wrong node. Expected %s, got %s", path, curr.name))
				}
			}

			continue
		}

		if strings.Contains(line, "$ ls") {
			// Not implemented
			continue
		}

		if strings.Contains(line, "dir ") {
			path := strings.Replace(line, "dir ", "", 1)
			// Append new dir to current children
			curr.children = append(curr.children, &File{
				name:     path,
				fileType: "dir",
				children: []*File{},
				size:     0,
				parent:   curr,
				depth:    curr.depth + 1,
			})
			continue
		}

		// Add child to tree
		parts := strings.Split(line, " ")
		size, _ := strconv.ParseInt(parts[0], 10, 32)
		curr.children = append(curr.children, &File{
			name:     parts[1],
			fileType: "file",
			size:     int(size),
			depth:    curr.depth + 1,
			parent:   curr,
		})
	}

	// Part1 Answer
	queue := []*File{root}
	sums := []uint64{}
	names := []string{}

	c := 0
	for len(queue) > 0 {
		current := queue[0] // Fetch
		queue = queue[1:]   // Pop

		if current.fileType == "dir" {
			// for _, c := range current.children {
			// 	fmt.Printf("Child of %s - %s - %d\n", current.name, c.name, c.size)
			// }
			s := Sum(0, current)
			// fmt.Printf("Summing %v, val: %d \n", current.name, s)
			// fmt.Println()
			sums = append(sums, s)
			names = append(names, current.name)
		}

		queue = append(queue, current.children...)
		// time.Sleep(1 * time.Second)
		c += 1
	}

	for _, v := range sums {
		if v <= PART_1_THRESHOLD {
			// fmt.Println(names[i], v)
			part1 += int(v)
		}
	}

	return internal.Answer{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func PrintDir(f *File) *File {
	fmt.Printf("%s %s\n", strings.Repeat("\t", f.depth), f.name)

	if f.fileType == "file" {
		return nil
	}

	for _, c := range f.children {
		s := PrintDir(c)

		if s != nil {
			return s
		}
	}

	return nil
}
