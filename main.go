package main

import (
	"advent-of-code/year2015"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting Year 2015 Solutions")

	start := time.Now()

	year2015.Day1()
	year2015.Day2()
	year2015.Day3()
	year2015.Day4()
	year2015.Day5()
	year2015.Day6()

	elapsed := time.Since(start)

	fmt.Printf("\n\nYear 2015 Complete. Elapsed time: %s\n", elapsed)
}
