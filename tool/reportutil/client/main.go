package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	ru "github.com/timpointer/golang-demo/tool/reportutil"
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
						cli.IntFlag{Name: "worknumber, w", Value: 10},
					},
					Action: func(c *cli.Context) error {
						fmt.Println("yaml ist rad", c.String("lang"), c.Bool("forever"))

						ctx := context.Background()

						stream := ru.Generator(ctx)

						worknumber := c.Int("worknumber")
						fanout := make([]<-chan interface{}, worknumber)

						for i := 0; i < worknumber; i++ {
							fanout[i] = ru.Multiply(ctx, ru.Add(ctx, stream, "out"))
						}

						pipeline := ru.FanIn(ctx, fanout...)
						for v := range pipeline {
							fmt.Println(v)
						}

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
