package year2021

import (
	"advent-of-code/pkg/data"
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	hasWon bool
	nums   [][]BingoNumber
}

type BingoNumber struct {
	number  int
	checked bool
}

// --- Day 4: Giant Squid ---
func Day4() {
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

	fmt.Printf("--- Day 4: Giant Squid --- { 1: %d, 2: %d }\n", part1, part2)
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

		fmt.Println(temp + "\n")
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
