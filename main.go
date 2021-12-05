package main

import (
	"advent-of-code/year2021"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting Year 2015 Solutions")

	start := time.Now()

	// year2015.Day1()
	// year2015.Day2()
	// year2015.Day3()
	// year2015.Day4()
	// year2015.Day5()
	// year2015.Day6()
	// // year2015.Day7()
	// year2015.Day8()

	year2021.Day1()
	year2021.Day2()
	year2021.Day3()
	year2021.Day4()

	elapsed := time.Since(start)

	fmt.Printf("\n\nYear 2015 Complete. Elapsed time: %s\n", elapsed)
}
