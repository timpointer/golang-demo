package main

import (
	"log"
	"time"
)

func main() {
	var food = make(chan int)
	go func() {
		getFood(food)
	}()

	tick := time.Tick(5 * time.Second)

out:
	for {
		select {
		case <-food:
			log.Println("get food")
			break out
		case <-tick:
			log.Println("alert")
			break out
		default:
		}
		time.Sleep(1 * time.Second)
	}
}

func getFood(food chan int) {
	time.Sleep(2 * time.Second)
	food <- 1
}
