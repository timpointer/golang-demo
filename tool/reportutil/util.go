package reportutil

import (
	"context"
	"fmt"
)

func PipeBridge(stream <-chan interface{}, list ...Pipe) PipeFunc {
	return func(in <-chan interface{}) <-chan interface{} {
		for _, v := range list {
			in = v.Do(stream)
		}
		return in
	}
}

type PipeFunc func(<-chan interface{}) <-chan interface{}

func (f PipeFunc) Do(in <-chan interface{}) <-chan interface{} {
	return f(in)
}

type Pipe interface {
	Do(inStream <-chan interface{}) <-chan interface{}
}

type Handler interface {
	Handle(in interface{}) interface{}
}

type UtilPipe struct {
	Ctx     context.Context
	Handler Handler
}

func (p *UtilPipe) Do(inStream <-chan interface{}) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for v := range OrDone(p.Ctx, inStream) {
			outStream <- p.Handler.Handle(v)
		}
	}()
	return outStream
}

type AddHandler struct {
	Add string
}

func (h AddHandler) Handle(in interface{}) interface{} {
	return fmt.Sprintf("%s%s", in, h.Add)
}

type MultiplyHandler struct {
}

func (h MultiplyHandler) Handle(in interface{}) interface{} {
	return fmt.Sprintf("%s%s", in, in)
}

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
