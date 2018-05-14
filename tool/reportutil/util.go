package reportutil

import (
	"context"
	"fmt"
)

func Add(ctx context.Context, inStream <-chan interface{}, add string) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for v := range OrDone(ctx, inStream) {
			outStream <- fmt.Sprintf("%s%s", v, add)
		}
	}()
	return outStream
}

func Multiply(ctx context.Context, inStream <-chan interface{}) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for v := range OrDone(ctx, inStream) {
			outStream <- fmt.Sprintf("%s%s", v, v)
		}
	}()
	return outStream
}
