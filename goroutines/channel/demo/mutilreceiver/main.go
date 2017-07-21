package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

func main() {
	task := make(chan int)
	var sw sync.WaitGroup

	go func() {
		for i := 0; i < 10; i++ {
			task <- i
		}
		close(task)
	}()

	for i := 0; i < 2; i++ {
		sw.Add(1)
		go func(i int) {
			defer sw.Done()
			for v := range task {
				log.Println(i, "v", v)
			}
		}(i)
	}
	sw.Wait()
	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
}
