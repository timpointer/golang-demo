package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:     "task",
			Aliases:  []string{"t"},
			Category: "task",
			Usage:    "task",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "lang, l",
							Value: "english",
							Usage: "Language for the greeting",
						},
						cli.BoolFlag{Name: "forever, forevvarr"},
					},
					Action: func(c *cli.Context) error {
						fmt.Println("yaml ist rad", c.String("lang"), c.Bool("forever"))
						return nil
					},
				},
			},
		},
		{
			Name:     "second",
			Aliases:  []string{"s"},
			Category: "second",
			Usage:    "second",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
