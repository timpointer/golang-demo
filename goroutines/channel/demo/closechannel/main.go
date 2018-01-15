package main

import (
	"fmt"
	"time"
)

func main() {
	task := make(chan int, 10)
	go func() {
		defer close(task)
		for i := 1; i <= 10; i++ {
			task <- i
		}
	}()

	time.Sleep(time.Second * 5)

	for i := range task {
		fmt.Println("i:", i)
	}
	fmt.Println("end:")
}
