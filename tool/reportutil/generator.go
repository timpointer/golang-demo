package reportutil

import (
	"bufio"
	"context"
	"os"
)

func Generator(ctx context.Context) <-chan interface{} {
	f := os.Stdin
	stream := make(chan interface{})

	go func() {
		defer close(stream)
		scanner := bufio.NewScanner(f)
		const max = 64 * 1024 * 4
		buff := make([]byte, max)
		scanner.Buffer(buff, max)

		for scanner.Scan() {
			text := scanner.Text()
			stream <- text
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()
	return stream
}
