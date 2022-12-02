package year2022

import (
	"advent-of-code/pkg/util"
	"fmt"
)

func Solve() {
	PrintIntro()
	// util.Time(util.PrintFunc(Day1))
	util.Time(util.PrintFunc(Day2))
}

func PrintIntro() {
	s := `Welcome to blah
`
	fmt.Println(s)
}
