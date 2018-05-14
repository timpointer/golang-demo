package reportutil

import (
	"context"
	"sync"
)

type PipeFunc func() <-chan interface{}

func (f PipeFunc) Do() <-chan interface{} {
	return f()
}

type Pipe interface {
	Do() <-chan interface{}
}

func FanOut(worknumber int, pipe Pipe) []<-chan interface{} {
	fanout := make([]<-chan interface{}, worknumber)

	for i := 0; i < worknumber; i++ {
		fanout[i] = pipe.Do()
	}
	return fanout
}

func FanIn(ctx context.Context, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
