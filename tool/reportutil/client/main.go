package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

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

						worknumber := c.Int("worknumber") //配置线程数量
						ctx := context.Background()

						// 执行超过100秒自动取消
						ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
						defer cancel()

						// 心跳器
						heartbeat := ru.NewHeartbeat(time.Second)

						// 生产者
						var gen ru.Generator = &ru.StdinGenerator{F: os.Stdin, Ctx: ctx, Heartbeat: heartbeat}
						stream := gen.Produce()

						// 检测心跳
						go func() {
							for v := range heartbeat.Output() {
								ru.PrintHeartbeat(v)
							}
						}()

						//适配pipe接口
						f := func(ctx context.Context, in <-chan interface{}) ru.PipeFunc {
							return func() <-chan interface{} {
								return ru.Multiply(ctx, ru.Add(ctx, in, "out"))
							}
						}(ctx, stream)

						//分配给多个管道执行
						fanout := ru.FanOut(worknumber, ru.PipeFunc(f))
						//聚合多个管道结果
						pipeline := ru.FanIn(ctx, fanout...)

						//一个管道分叉成两个复制，两个管道同步消费
						p1, p2 := ru.Tee(ctx, pipeline)
						for v := range p2 {
							fmt.Println(v, "+", <-p1)
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
