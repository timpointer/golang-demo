package reportutil

import (
	"bufio"
	"context"
	"os"

	"golang.org/x/time/rate"

	log "github.com/sirupsen/logrus"
)

type Generator interface {
	Produce() <-chan interface{}
}

type StdinGenerator struct {
	F           *os.File
	Ctx         context.Context
	Heartbeat   Heartbeat
	RateLimiter *rate.Limiter
}

func (g *StdinGenerator) Produce() <-chan interface{} {
	stream := make(chan interface{})

	go func() {
		defer close(stream)
		defer g.Heartbeat.Close()

		scanner := bufio.NewScanner(g.F)
		const max = 64 * 1024 * 4
		buff := make([]byte, max)
		scanner.Buffer(buff, max)

		for scanner.Scan() {
			text := scanner.Text()
			if err := g.RateLimiter.Wait(g.Ctx); err != nil {
				log.WithFields(log.Fields{
					"err": err,
				}).Error("g.rateLimiter.Wait")
				return
			}
			g.Heartbeat.SendPluse()

			select {
			case stream <- text:
				g.Heartbeat.Add()
			case <-g.Ctx.Done():
				return
			}
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()
	return stream
}
