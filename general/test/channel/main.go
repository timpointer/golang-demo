package main

import "log"

func main() {
}

type chanBuffer struct {
	input chan int
	ouput chan int
	close chan bool
}

func newChanBuffer(in, out chan int) *chanBuffer {
	return &chanBuffer{
		input: in,
		ouput: out,
		close: make(chan bool, 1),
	}
}

func (b *chanBuffer) Close() {
	b.close <- true
}

func (b *chanBuffer) Run() {
	go func() {
		select {
		case t := <-b.input:
			log.Println("t:", t)
		case <-b.close:
			for range b.input {
			}
			return
		}
	}()
}
