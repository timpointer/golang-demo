package main

import (
	"fmt"
	"log"
	"time"
)

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	natuals := &natuals{make(chan int), make(chan int)}
	counter := &counter{natuals.Out(), make(chan int)}
	go natuals.natual()
	go counter.count()
	go func() {
		time.Sleep(1 * time.Second)
		natuals.Close()
	}()
	printer(counter.Out())
}

type natuals struct {
	out   chan int
	close chan int
}

func (n *natuals) Out() chan int {
	return n.out
}

func (n *natuals) Close() {
	close(n.close)
}

func (n *natuals) natual() {
	for x := 0; x < 100; x++ {
		select {
		case n.out <- x:
			log.Println("send ", x)
			time.Sleep(100 * time.Millisecond)
		case <-n.close:
			log.Println("close")
			close(n.out)
			return
		}
	}
	close(n.out)
}

type counter struct {
	in  chan int
	out chan int
}

func (n *counter) count() {
	for x := range n.in {
		n.out <- x
	}
	close(n.out)
}

func (n *counter) Out() chan int {
	return n.out
}
