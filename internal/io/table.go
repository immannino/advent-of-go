package io

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func Print(r [][]string) {
	headerFmt := color.New(color.FgHiMagenta, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgCyan).SprintfFunc()

	tbl := table.New("Puzzle", "Part 1", "Part 2", "Time")
	tbl.SetRows(r)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	tbl.Print()
}
