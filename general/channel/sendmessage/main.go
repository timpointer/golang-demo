package main

import (
	"log"
	"time"
)

func main() {
	var task = make(chan int)
	var task2 = make(chan int)
	go func() {
		t := <-task2
		log.Println("t", t)
	}()
	go func() {
		task2 <- <-task
	}()
	task <- 1

	time.Sleep(time.Second)
	log.Println("end")
}
