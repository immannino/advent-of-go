package year2022

import (
	"advent-of-code/pkg/util"
	"fmt"
)

func Solve() {
	PrintIntro()
	fmt.Println(util.Time(Day1))
	fmt.Println(util.Time(Day2))
	fmt.Println(util.Time(Day3))
	fmt.Println(util.Time(Day4))
}

func PrintIntro() {
	s := `Welcome to blah
`
	fmt.Println(s)
}
