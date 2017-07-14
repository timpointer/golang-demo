package main

import (
	"log"
	"sync"
)

type workers struct {
	wg       sync.WaitGroup
	itemChan chan interface{}
	errChan  chan error
	number   int
	flist    []func(chan interface{}, chan error)
}

func newWorkers(flist []func(chan interface{}, chan error)) *workers {
	return &workers{
		itemChan: make(chan interface{}),
		errChan:  make(chan error, len(flist)),
		number:   len(flist),
		flist:    flist,
	}
}

func (w *workers) run() ([]interface{}, error) {
	log.Println("w", w.number, len(w.flist))
	for i := 0; i < w.number; i++ {
		w.wg.Add(1)
		go func(i int) {
			defer w.wg.Done()
			w.flist[i](w.itemChan, w.errChan)
		}(i)
	}
	go func() {
		w.wg.Wait()
		close(w.itemChan)
		close(w.errChan)
	}()

	var list []interface{}
	for item := range w.itemChan {
		list = append(list, item)
	}

	for err := range w.errChan {
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

func main() {
	var list []int
	var listname = []string{"sdf", "tim"}
	var flist = []func(chan interface{}, chan error){}
	for _, name := range listname {
		name := name
		f := func() func(chan interface{}, chan error) {
			return func(i chan interface{}, e chan error) {
				log.Println("name", name)
				i <- 1
			}
		}()
		flist = append(flist, f)
	}

	workers := newWorkers(flist)
	result, err := workers.run()
	if err != nil {
		log.Fatal("error")
	}

	for _, item := range result {
		list = append(list, item.(int))
	}
	log.Println("list", list)
}
