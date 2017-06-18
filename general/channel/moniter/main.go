package main

import "fmt"

func main() {
	first := make(chan int)
	second := make(chan int)
	go create(first)
	go worker(first, second)
	print(second)
}

func create(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func worker(in <-chan int, out chan<- int) {
	for item := range in {
		out <- item * 2
	}
	close(out)
}

func print(in <-chan int) {
	for item := range in {
		fmt.Println("item:", item)
	}
}
