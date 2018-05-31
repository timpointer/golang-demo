package main

import "time"

func main() {
	dataStream := make(chan interface{})
	go func() {
		defer close(dataStream)
		time.Sleep(time.Second)

	}()
	<-dataStream
}
