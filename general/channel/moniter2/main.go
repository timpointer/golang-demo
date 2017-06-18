package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var status = make(map[int]int)

var tens = 10 * time.Second
var one = 1 * time.Second

func main() {
	var sw sync.WaitGroup
	msg := make(chan int)
	go readStatus()
	go moniterStatus()
	go collectionStatus(msg)
	for i := 1; i <= 3; i++ {
		sw.Add(1)
		go longtask(i, &sw, msg)
	}
	sw.Wait()
	close(msg)
	fmt.Println("program finish")
}

func collectionStatus(in <-chan int) {
	for {
		key := <-in
		status[key]++
	}
}

func moniterStatus() {
	c := time.Tick(time.Second)
	for {
		<-c
		printStatus()
	}
}

func printStatus() {
	for key, value := range status {
		var pg string
		for i := 0; i <= value; i++ {
			pg += "*"
		}
		fmt.Printf("task %d ,progress %s\n", key, pg)
	}
}

func readStatus() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		input.Text()
		printStatus()
	}
}

func longtask(key int, sw *sync.WaitGroup, out chan<- int) {
	defer sw.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(one)
		out <- key
	}
}
