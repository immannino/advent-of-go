package cmd

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	dayFlags = []cli.Flag{
		&cli.IntFlag{
			Name:    "day",
			Usage:   "Run a specific day",
			Aliases: []string{"d"},
		},
	}
)

func Start() {
	app := &cli.App{
		Name:  "aoc",
		Usage: "Advent of Code solutions",
		Action: func(*cli.Context) error {
			NewYear2015().SolvePretty()
			NewYear2021().SolvePretty()
			NewYear2022().SolvePretty()
			NewYear2023().SolvePretty()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "2015",
				Flags: dayFlags,
				Action: func(ctx *cli.Context) error {
					d := NewYear2015()
					if ctx.Int("day") > 0 {
						return d.SolveSingle(ctx.Int("day"))
					}

					d.SolvePretty()
					return nil
				},
			},
			{
				Name:  "2021",
				Flags: dayFlags,
				Action: func(ctx *cli.Context) error {
					d := NewYear2021()
					if ctx.Int("day") > 0 {
						return d.SolveSingle(ctx.Int("day"))
					}

					d.SolvePretty()
					return nil
				},
			},
			{
				Name:  "2022",
				Flags: dayFlags,
				Action: func(ctx *cli.Context) error {
					d := NewYear2022()
					if ctx.Int("day") > 0 {
						return d.SolveSingle(ctx.Int("day"))
					}

					d.SolvePretty()
					return nil
				},
			},
			{
				Name:  "2023",
				Flags: dayFlags,
				Action: func(ctx *cli.Context) error {
					d := NewYear2023()
					if ctx.Int("day") > 0 {
						return d.SolveSingle(ctx.Int("day"))
					}

					d.SolvePretty()
					return nil
				},
			},
			{
				Name:   "readme",
				Action: GenerateReadme,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

const readmeTemplate = `# Advent of Go

> Golang solutions to Advent of Code

## Completed Solutions

{{range $val := .}}

{{ $val }}

{{end}}

## Past AoC Projects

[aoc.spaghet.me](https://aoc.spaghet.me) - 2022 Solutions using Web Workers [github](https://github.com/immannino/aoc-2020)

[Original attempts 2017-2018](https://github.com/immannino/AdventOfCode)`

func GenerateReadme(ctx *cli.Context) error {
	solutions := []string{
		NewYear2023().SolvePrettyToString(),
		NewYear2022().SolvePrettyToString(),
		NewYear2021().SolvePrettyToString(),
		NewYear2015().SolvePrettyToString(),
	}

	t := template.Must(template.New("").Parse(readmeTemplate))

	var b bytes.Buffer
	err := t.Execute(&b, solutions)
	if err != nil {
		log.Println("error executing template, " + err.Error())
		return err
	}

	err = ioutil.WriteFile("./README.md", b.Bytes(), 0644)
	if err != nil {
		log.Println("error creating README.md, " + err.Error())
		return err
	}

	log.Println("Readme Successfully Created")
	return nil
}
