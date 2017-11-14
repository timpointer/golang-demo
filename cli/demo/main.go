package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func init() {
}

func main() {

	var language string
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang,l",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
			EnvVar:      "APP_LANG,LANG",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q", c.Args().Get(0))
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "import",
			Usage: "pipetaskimport",
			Action: func(c *cli.Context) error {
				fmt.Println("pipetaskimport", language)
				fmt.Println("pipe", language)
				return nil
			},
		},
		{
			Name:  "load",
			Usage: "pipetaskload",
			Action: func(c *cli.Context) error {
				fmt.Println("pipetaskload")
				return nil
			},
		},
		{
			Name:  "fresh",
			Usage: "pipetaskrefresh",
			Action: func(c *cli.Context) error {
				fmt.Println("pipetaskrefresh")
				return nil
			},
		},
		{
			Name:    "pipe",
			Aliases: []string{"p"},
			Usage:   "pipeline",
			Action: func(c *cli.Context) error {
				fmt.Println("pipeline")
				return nil
			},
		},
		{
			Name:  "config",
			Usage: "show config information",
			Action: func(c *cli.Context) error {
				fmt.Println("config")
				return nil
			},
		},
	}

	app.Name = "overviewloader"
	app.Usage = "create overview report"
	app.Version = "0.0.1"
	sort.Sort(cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
