package main

import (
	"log"
	"strconv"

	"github.com/timpointer/golang-demo/general/test/channel5/worker"
)

func main() {
	tasks := make(chan int)

	for i := 0; i < 10; i++ {
		wer := worker.NewMyWorker("stage1name" + strconv.Itoa(i))
		wer.Input(tasks)
		wer.Run()
	}

	wer2 := worker.NewMyWorker("worker1")
	wer2.Input(wer.Output())
	wer2.Run()

	go func() {
		for i := 0; i < 10; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	for v := range wer2.Output() {
		log.Println("receive", v)
	}
	log.Println("end")
}
