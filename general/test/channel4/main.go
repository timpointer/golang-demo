package main

import (
	"log"
	"time"
)

func main() {
	done := make(chan bool)
	finish := make(chan bool)
	fileSizes := make(chan int)
	go func() {
		for {
			select {
			case <-done:
				// Drain fileSizes to allow existing goroutines to finish.
				for range fileSizes {
					// Do nothing.

					time.Sleep(4 * time.Second)
					log.Println("delete task")
				}
				finish <- true
				return
			case size, ok := <-fileSizes:
				time.Sleep(4 * time.Second)
				log.Println("size", size, "ok", ok)
			}
		}
	}()
	go func() { fileSizes <- 1 }()
	go func() { fileSizes <- 2 }()
	go func() {
		done <- true
		close(fileSizes)
	}()
	go func() { fileSizes <- 2 }()
	go func() { fileSizes <- 2 }()
	<-finish
}
