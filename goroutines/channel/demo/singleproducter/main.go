package main

import "log"

func main() {
	task := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			task <- i
		}
		close(task)
	}()

	for v := range task {
		log.Println("v", v)
	}
}
