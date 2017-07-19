package main

import (
	"log"
	"net/http"
	"time"
)

type database struct {
	task chan int
}

func main() {
	db := &database{task: make(chan int, 2)}

	var start = make(chan int)
	var todotask = make(chan int)

	// long task
	go func() {
		time.Sleep(5 * time.Second)
		start <- 1
	}()

	var buf []int
	go func() {
	stage1:
		for {
			select {
			case t := <-db.task:
				buf = append(buf, t)
			case <-start:
				break stage1
			}
		}
		log.Println("stage2")
		go func() {
			for _, t := range buf {
				todotask <- t
			}
		}()

		for {
			todotask <- <-db.task
		}
	}()

	go func() {
		for t := range todotask {
			log.Println("t:", t)
		}
	}()

	http.HandleFunc("/stream", db.stream)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println(err)
	}
}

var count int

func (d *database) stream(w http.ResponseWriter, r *http.Request) {
	d.task <- count
	log.Println("success", count)
	count++
	w.Write([]byte("success"))
}
