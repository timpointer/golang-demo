package main

import (
	"log"
	"sync"
)

func main() {
	task := make(chan int)
	workers := NewWorker(3, task, func(id, i int) {
		log.Println(id, "i", i)
	})
	workers.Run()
	go func() {
		for i := 0; i < 10; i++ {
			task <- i
		}
		close(task)
	}()
	workers.Wait()
}

type Workers struct {
	input    chan int
	sw       sync.WaitGroup
	number   int
	cousumer func(int, int)
}

func NewWorker(num int, input chan int, f func(int, int)) *Workers {
	return &Workers{number: num, input: input, cousumer: f}
}

func (w *Workers) Wait() {
	w.sw.Wait()
}
func (w *Workers) Run() {
	for i := 0; i < w.number; i++ {
		w.sw.Add(1)
		go func(i int) {
			defer w.sw.Done()
			for v := range w.input {
				w.cousumer(i, v)
			}
		}(i)
	}
}

type multiSaver struct {
	taskChanMap map[string]chan int
	mutex       sync.Mutex
	wg          sync.WaitGroup
}

func newMultiSaver() *multiSaver {
	return &multiSaver{taskChanMap: make(map[string]chan int)}
}

func (w *multiSaver) close() {
	for _, ch := range w.taskChanMap {
		close(ch)
	}
	w.wg.Wait()
}

func (w *multiSaver) getInput(name string, f func(int, int)) chan int {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	taskChan, ok := w.taskChanMap[name]
	if ok == false {
		taskChan = make(chan int)
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()

			log.Println(name, " saver is closed")
		}()
		w.taskChanMap[name] = taskChan
	}
	return taskChan
}
