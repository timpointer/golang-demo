package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create("cpu-profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	fmt.Println(action(40))
	wait()
	fmt.Println(fibonacci(45))
	pprof.StopCPUProfile()
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func wait() {
	time.Sleep(time.Second)
}

func action(n int) int {
	if n < 2 {
		return n
	}
	return action(n-1) + action(n-2)
}
