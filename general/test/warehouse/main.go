package main

import (
	"log"
	"net/http"
	"time"
)

type database struct {
	task     chan int
	receiver Receiver
}

func main() {
	loader := &Loader{make(chan int)}
	loader.run()

	cb := newChanBuffer(loader)
	cb.Run()

	db := &database{task: make(chan int, 2), receiver: cb}

	// long task
	go longTask(cb)

	http.HandleFunc("/stream", db.stream)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println(err)
	}
	panic("wer")
}

type Loader struct {
	keyChan chan int
}

func (c *Loader) get() chan int {
	return c.keyChan
}

func (c *Loader) receive(i int) {
	c.keyChan <- i
}

func (c *Loader) run() {
	go func() {
		log.Println("todotask")
		for t := range c.keyChan {
			log.Println("t:", t)
		}
	}()
}

var count int

func (d *database) stream(w http.ResponseWriter, r *http.Request) {
	d.receiver.receive(count)
	log.Println("success", count)
	count++
	w.Write([]byte("success"))
}

func longTask(start Starter) {
	time.Sleep(5 * time.Second)
	start.start()
}

type Receiver interface {
	receive(int)
}

type Sender interface {
	send(int)
}

type Starter interface {
	start()
}

type chanBuffer struct {
	buf      []int
	done     chan bool
	close    chan bool
	input    chan int
	receiver Receiver
}

func newChanBuffer(receiver Receiver) *chanBuffer {
	return &chanBuffer{
		input:    make(chan int),
		receiver: receiver,
		done:     make(chan bool, 1),
		close:    make(chan bool, 1),
	}
}

func (b *chanBuffer) start() {
	b.done <- true
}

func (b *chanBuffer) Close() {
	b.close <- true
}

func (b *chanBuffer) receive(i int) {
	b.input <- i
}

func (b *chanBuffer) Run() {
	go func() {
	stage1:
		for {
			select {
			case t := <-b.input:
				b.buf = append(b.buf, t)
			case <-b.done:
				break stage1
			case <-b.close:
				for range b.input {
				}
				return
			}
		}
		log.Println("stage2")
		go func() {
			for _, t := range b.buf {
				b.receiver.receive(t)
			}
		}()

		for {
			b.receiver.receive(<-b.input)
		}
	}()
}
