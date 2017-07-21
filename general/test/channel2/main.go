package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	loader := &Loader{make(chan int)}
	loader.Run()

	cb := NewChanBuffer(loader)
	cb.Run()

	go func(beginer Beginer) {
		time.Sleep(5 * time.Second)
		beginer.Begin()
	}(cb)
	go func() {
		cb.Close()
		loader.Close()
	}()
	c := time.Tick(2 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

type Loader struct {
	i chan int
}

func (l *Loader) Receive(i int) bool {
	l.i <- i
	return true
}
func (l *Loader) Close() {
	close(l.i)
}

func (l *Loader) Run() {
	go func() {
		for i := range l.i {

			log.Println("i", i)
		}
	}()
}

type ReceiveWorker interface {
	Receiver
	Worker
}

type Worker interface {
	Run()
	Close()
}

type Receiver interface {
	Receive(int) bool
}

type Sender interface {
	Send(int)
}

type Beginer interface {
	Begin()
}

type ChanBuffer struct {
	buf      []int
	begin    chan bool
	close    chan bool
	input    chan int
	receiver Receiver
}

func NewChanBuffer(receiver Receiver) *ChanBuffer {
	return &ChanBuffer{
		input:    make(chan int),
		receiver: receiver,
		begin:    make(chan bool, 1),
		close:    make(chan bool, 1),
	}
}

func (b *ChanBuffer) Begin() {
	b.begin <- true
}

func (b *ChanBuffer) Close() {
	b.close <- true
	close(b.input)
}

func (b *ChanBuffer) Receive(i int) bool {
	b.input <- i
	return true
}

func (b *ChanBuffer) Run() {
	go func() {
	stage1:
		for {
			select {
			case t := <-b.input:
				b.buf = append(b.buf, t)
			case <-b.begin:
				break stage1
			case <-b.close:
				for range b.input {
				}
				log.Println("b.close")
				return
			}
		}
		log.Println("stage2")
		go func() {
			for _, t := range b.buf {
				b.receiver.Receive(t)
			}
		}()

		for {
			b.receiver.Receive(<-b.input)
		}
	}()
}
