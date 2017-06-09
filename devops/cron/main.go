package main

import "fmt"
import "github.com/robfig/cron"
import "time"

func main() {
	c := cron.New()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()

	// Funcs are invoked in their own goroutine, asynchronously.

	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })

	// Inspect the cron job entries' next and previous run times.
	inspect(c.Entries())
	time.Sleep(time.Duration(60 * time.Second))
	c.Stop() // Stop the scheduler (does not stop any jobs already running).
}

func inspect(entries []*cron.Entry) {
	for _, entry := range entries {
		fmt.Println(entry)
	}
}
