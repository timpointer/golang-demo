package main

import (
	"log"
	"time"

	"github.com/timpointer/golang-demo/general/test/mycron"
)

func main() {
	cron := mycron.New()
	cron.AddJob("* * * * * *", &task{"tim"})
	cron.AddJob("* * * * * *", &task{"mary"})
	cron.Start()
	time.Sleep(2 * time.Second)
	cron.Remove("tim")
	time.Sleep(5 * time.Second)
	cron.Stop()
	log.Println("stop")
	time.Sleep(1 * time.Second)
	log.Println("start")
	cron.AddJob("* * * * * *", &task{"gek"})
	cron.Start()
	time.Sleep(1000 * time.Hour)
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
