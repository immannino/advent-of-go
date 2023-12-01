package internal

import (
	"advent-of-code/internal/io"
	"errors"
	"fmt"
	"time"
)

type Puzzle struct {
	ID         int
	Name       string
	PuzzleFunc func() (string, string)
}

type PuzzleInterface interface {
	GetID() int
	Run() (part1 string, part2 string)
	RunWithTable() []string
}

func NewPuzzle(ID int, Name string, f func() (string, string)) PuzzleInterface {
	return Puzzle{
		ID:         ID,
		Name:       Name,
		PuzzleFunc: f,
	}
}

func (p Puzzle) GetID() int {
	return p.ID
}

func (p Puzzle) Run() (part1 string, part2 string) {
	return p.PuzzleFunc()
}

func (p Puzzle) RunWithTable() []string {
	start := time.Now()
	p1, p2 := p.PuzzleFunc()
	end := time.Since(start)
	return []string{p.Name, p1, p2, fmt.Sprintf("%v", end)}
}

type Year struct {
	Title   string
	Puzzles []PuzzleInterface
}

type YearInterface interface {
	Solve()
	SolveSingle(id int) error
	SolvePretty()
}

func (y Year) SolveSingle(id int) error {
	for _, v := range y.Puzzles {
		if v.GetID() == id {
			p1, p2 := v.Run()
			fmt.Println(y.Title)
			fmt.Println(fmt.Sprintf("Day %d:", id), p1, p2)
			return nil
		}
	}

	return errors.New("No puzzle by that ID")
}

func (y Year) Solve() {
	start := time.Now()
	for _, d := range y.Puzzles {
		d.Run()
	}
	end := time.Since(start)
	fmt.Printf("%s: Total Time: %v\n", y.Title, end)
}

func (y Year) SolvePretty() {
	start := time.Now()
	r := [][]string{{y.Title, "", "", ""}}

	for _, d := range y.Puzzles {
		r = append(r, d.RunWithTable())
	}

	end := time.Since(start)
	r = append(r, []string{"Total", "", "", fmt.Sprintf("%v", end)})
	io.Print(r)
}

func (y Year) SolvePrettyToString() string {
	start := time.Now()
	r := [][]string{{y.Title, "", "", ""}}

	for _, d := range y.Puzzles {
		r = append(r, d.RunWithTable())
	}

	end := time.Since(start)
	r = append(r, []string{"Total", "", "", fmt.Sprintf("%v", end)})

	return io.ToString(r)
}
