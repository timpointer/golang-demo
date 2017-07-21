package worker

import "log"

type Producter interface {
	Output() chan int
}

type Consumer interface {
	Input(chan int)
}

type Worker interface {
	Producter
	Consumer
	Run()
}

type MyWorker struct {
	task chan int
	toDo chan int
	name string
}

func NewMyWorker(name string) *MyWorker {
	return &MyWorker{task: make(chan int), name: name}
}

func (w *MyWorker) Output() chan int {
	return w.task
}

func (w *MyWorker) Input(i chan int) {
	w.toDo = i
}

func (w *MyWorker) Run() {
	go func() {
		for do := range w.toDo {
			w.task <- do
			log.Println("do the work", do)
		}
		close(w.task)
	}()
}
