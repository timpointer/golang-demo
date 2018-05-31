package reportutil

import (
	"context"
	"sync"
)

func FanOutIn(ctx context.Context, worknumber int, in <-chan interface{}, pipe Pipe) <-chan interface{} {
	fanout := FanOut(worknumber, in, pipe)

	//聚合多个管道结果
	return FanIn(ctx, fanout...)
}

func FanOut(worknumber int, in <-chan interface{}, pipe Pipe) []<-chan interface{} {
	fanout := make([]<-chan interface{}, worknumber)

	for i := 0; i < worknumber; i++ {
		fanout[i] = pipe.Do(in)
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
