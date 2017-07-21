package main

import (
	"log"

	"gopkg.in/robfig/cron.v2"
)

func main() {
	cron := cron.New()
	go func() {
		//	cron.AddJob("* * * * * *", &task{"tim"})
	}()
	cron.AddJob("* * * * * *", &task{"tim"})
	cron.AddJob("* * * * * *", &task{"mary"})
}

type task struct {
	name string
}

func (t *task) Run() {
	log.Println(t.name, " task run")
}

func (t *task) ID() string {
	return t.name
}
