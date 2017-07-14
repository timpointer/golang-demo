package main

import (
	"log"
	"time"
)

func main() {
	for {
		select {
		case <-time.Tick(time.Second):
			log.Println("tick")
		default:
			log.Println("second 1")
			time.Sleep(2 * time.Second)
		}
	}
}
