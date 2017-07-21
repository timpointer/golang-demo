package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}
func queryAll() int {
	ch := make(chan int)
	go func() { ch <- query() }()
	go func() { ch <- query() }()
	go func() { ch <- query() }()
	return <-ch
}
func main() {
	var ch chan int
	if false {
		ch = make(chan int, 1)
		ch <- 1
	}
	go func(ch chan int) {
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
