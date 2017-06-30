package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var done = make(chan struct{})

func isCancel() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	var paths []string
	if len(os.Args) < 1 {
		paths[0] = "."
	} else {
		paths = os.Args[1:]
	}
	go func() {
		var by = make([]byte, 1)
		os.Stdin.Read(by)
		log.Printf("close program %v\n", by)
		close(done)
	}()
	var size = make(chan int64)
	var wg sync.WaitGroup
	for _, path := range paths {
		wg.Add(1)
		go search(&wg, path, size)
	}

	go func() {
		wg.Wait()
		close(size)
	}()

	var count int
	var totalsize int64
loop:
	for {
		select {
		case s, ok := <-size:
			if ok != true {
				break loop
			}
			count++
			totalsize += s
		case <-done:
			return
		}
	}
	fmt.Printf("%d files, %d totalsize\n", count, totalsize)
}

var flag = make(chan struct{}, 2)

func search(wg *sync.WaitGroup, path string, size chan<- int64) {
	flag <- struct{}{}
	defer func() { <-flag }()
	defer wg.Done()
	if isCancel() {
		return
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("file ", err)
	}
	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			go search(wg, filepath.Join(path, file.Name()), size)
		} else {
			size <- file.Size()
		}
	}
	return
}
