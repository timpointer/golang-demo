package main

func main() {

	ch := channel(make(chan int))

	go ch.read()
	ch.send(2)

}

type channel chan int

func (c *channel) send(input int) {
	*c <- input
}

func (c *channel) read() int {
	return <-*c
}
