package main

import "log"
import "sync"

func main() {
	task := make(chan int)
	var sw sync.WaitGroup

	for i := 0; i < 2; i++ {
		sw.Add(1)
		go func() {
			defer sw.Done()
			for i := 0; i < 10; i++ {
				task <- i
			}
		}()
	}

	go func() {
		sw.Wait()
		close(task)
	}()

	for v := range task {
		log.Println("v", v)
	}
}
