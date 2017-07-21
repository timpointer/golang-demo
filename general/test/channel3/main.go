package main

import (
	"fmt"
	"runtime"
)

func main() {
	done := make(chan bool)
	worker := new(done)
	worker.run()

	for j := 1; j <= 3; j++ {
		worker.receive(j)
		fmt.Println("sent job", j)
	}
	worker.close()
	fmt.Println("sent all jobs")
	<-done
	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
}

type worker struct {
	jobs chan int
	done chan bool
}

func new(done chan bool) *worker {
	return &worker{make(chan int, 5), done}
}

func (w *worker) receive(i int) {
	w.jobs <- i
}

func (w *worker) close() {
	close(w.jobs)
}

func (w *worker) run() {
	go func() {
		for {
			j, more := <-w.jobs
			if more {
				fmt.Println("received", j)
			} else {
				fmt.Println("received all jobs")
				w.done <- true
				return
			}
		}
	}()

}
