package main

import "log"

func main() {
	var c1, c2 <-chan interface{}
	var c3 = make(chan<- interface{}, 1)

	select {
	case <-c1:
	case <-c2:
	case c3 <- struct{}{}:
		log.Println("dsdfsdf")
	}
}
