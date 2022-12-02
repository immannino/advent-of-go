package year2022

import (
	"advent-of-code/pkg/data"
	"fmt"
	"strings"
)

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

func Day2() string {
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

	return fmt.Sprintf("---| Day 2 Rock Paper Scissors - 1: %d 2: %d |---\n", part1, part2)
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
