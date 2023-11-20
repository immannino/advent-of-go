package cmd

import (
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
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
