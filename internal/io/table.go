package io

import (
	"bytes"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
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

func ToString(r [][]string) string {
	var b bytes.Buffer
	table := tablewriter.NewWriter(&b)

	table.SetHeader([]string{"Puzzle", "Part 1", "Part 2", "Time"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(r)
	table.Render()

	return b.String()
}
